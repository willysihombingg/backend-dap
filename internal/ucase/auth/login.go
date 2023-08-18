package auth

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"

	"gitlab.com/willysihombing/task-c3/internal/appctx"
	"gitlab.com/willysihombing/task-c3/internal/connector/supabase/user"
	"gitlab.com/willysihombing/task-c3/internal/consts"
	"gitlab.com/willysihombing/task-c3/internal/middleware"
	"gitlab.com/willysihombing/task-c3/internal/presentations"
	"gitlab.com/willysihombing/task-c3/internal/ucase/contract"
	"gitlab.com/willysihombing/task-c3/pkg/logger"
	"gitlab.com/willysihombing/task-c3/pkg/tracer"
)

type loginUser struct {
	connectorUser user.Connectorer
}

func NewLoginUser(connectorUser user.Connectorer) contract.UseCase {
	return &loginUser{connectorUser: connectorUser}
}

// Serve partner list data
func (u *loginUser) Serve(dctx *appctx.Data) appctx.Response {
	var (
		ctx   = tracer.SpanStart(dctx.Request.Context(), "ucase.document_stamp_category")
		param presentations.LoginUser
		err   = dctx.Cast(&param)
		lf    = logger.NewFields(
			logger.EventName("CreateNewActivity"),
			logger.Any("query", fmt.Sprintf("*")),
		)
		response []user.ResponseGetUser
	)

	request := u.connectorUser.GetByEmail(param.Email)
	err = u.connectorUser.Send(ctx, request, &response)

	if err != nil {
		tracer.SpanError(ctx, err)
		logger.ErrorWithContext(ctx, fmt.Sprintf("all users is not connector erorr", err), lf...)
		return *appctx.NewResponse().WithCode(consts.CodeInternalServerError).WithMessage("Get All Users is empty")
	}

	if len(response) == 0 {
		tracer.SpanError(ctx, err)
		logger.ErrorWithContext(ctx, fmt.Sprintf("no users"), lf...)
		return *appctx.NewResponse().WithCode(consts.CodeNotFound).WithMessage("Users not found")
	}

	// Comparing the password with the hash
	err = bcrypt.CompareHashAndPassword([]byte(response[0].Password), []byte(param.Password))

	if err != nil {
		tracer.SpanError(ctx, err)
		logger.ErrorWithContext(ctx, fmt.Sprintf("all user is not connector erorr", err), lf...)
		return *appctx.NewResponse().WithCode(consts.CodeAuthenticationFailure).WithMessage("Wrong Password")
	}

	token, err := middleware.GenerateToken(response[0].ID, response[0].Email, response[0].Role, "jasonvonpeng18", 2)
	// fmt.Println("CEK RESPONSE :", response[0])
	data := presentations.LoginResponse{
		Token:      token,
		UserId:     response[0].ID,
		Name:       response[0].Name,
		Email:      response[0].Email,
		Role:       response[0].Role,
		Created_at: response[0].CreatedAt,
	}

	if err != nil {
		tracer.SpanError(ctx, err)
		logger.ErrorWithContext(ctx, fmt.Sprintf("generate token err", err), lf...)
		return *appctx.NewResponse().WithCode(consts.CodeInternalServerError).WithMessage("Fail generate tokne")
	}

	return *appctx.NewResponse().WithCode(consts.CodeSuccess).WithData(data).WithMessage("GetDetail Success")
}
