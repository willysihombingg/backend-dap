package middleware

import (
	"context"
	"fmt"
	"net/http"

	"github.com/kataras/jwt"
	"gitlab.com/willysihombing/task-c3/internal/appctx"
	"gitlab.com/willysihombing/task-c3/internal/connector/supabase/user"
	"gitlab.com/willysihombing/task-c3/internal/consts"
	"gitlab.com/willysihombing/task-c3/pkg/logger"
	"gitlab.com/willysihombing/task-c3/pkg/tracer"
)

type validateBearer struct {
	connectorUser user.Connectorer
}

func NewValidateBearer(connectorUser user.Connectorer) *validateBearer {
	return &validateBearer{connectorUser: connectorUser}
}

func (m *validateBearer) ValidateToken(r *http.Request, _ *appctx.Config) int {
	var (
		response []user.ResponseGetUser
		lf       = logger.NewFields(
			logger.EventName("MiddlewareMerchant"),
		)
	)

	ctx := tracer.SpanStart(r.Context(), "middleware.merchant_auth")
	defer tracer.SpanFinish(ctx)

	type MyClaims struct {
		ID       uint64
		Email    string
		Role     string
		Expiry   int64
		IssuedAt int
		Issuer   string
	}

	token, err := ParseHeaderAuthToken("Bearer", r)
	// fmt.Println("TOKEN: ", token)

	clm := MyClaims{}
	verifiedToken, err := jwt.Verify(jwt.HS256, []byte("jasonvonpeng18"), []byte(token))

	if err != nil {
		return consts.CodeAuthenticationFailure
	}

	err = verifiedToken.Claims(&clm)

	if err != nil {
		tracer.SpanError(ctx, err)
		logger.WarnWithContext(ctx, fmt.Sprintf("forbidden access got error %v", err), lf...)
		return consts.CodeAuthenticationFailure
	}

	request := m.connectorUser.GetByEmail(clm.Email)
	err = m.connectorUser.Send(ctx, request, &response)

	if err != nil {
		tracer.SpanError(ctx, err)
		logger.WarnWithContext(ctx, fmt.Sprintf("forbidden access got error %v", err), lf...)
		return consts.CodeAuthenticationFailure
	}

	if len(response) == 0 {
		fmt.Println(len(response))
		tracer.SpanError(ctx, err)
		logger.WarnWithContext(ctx, fmt.Sprintf("forbidden access got error %v", err), lf...)
		return consts.CodeAuthenticationFailure
	}

	req := r.WithContext(context.WithValue(r.Context(), consts.CtxUserInfo, response[0]))
	*r = *req
	return consts.CodeSuccess
}
