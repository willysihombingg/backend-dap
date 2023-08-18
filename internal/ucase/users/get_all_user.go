package user

import (
	"fmt"

	"gitlab.com/willysihombing/task-c3/internal/helper"

	"gitlab.com/willysihombing/task-c3/internal/appctx"
	"gitlab.com/willysihombing/task-c3/internal/connector/supabase/user"
	"gitlab.com/willysihombing/task-c3/internal/consts"
	"gitlab.com/willysihombing/task-c3/internal/ucase/contract"
	"gitlab.com/willysihombing/task-c3/pkg/logger"
	"gitlab.com/willysihombing/task-c3/pkg/tracer"
)

type getAllUser struct {
	connectorAdmin user.Connectorer
}

func GetAllUser(connectorAdmin user.Connectorer) contract.UseCase {
	return &getAllUser{
		connectorAdmin: connectorAdmin,
	}
}

func (u *getAllUser) Serve(dctx *appctx.Data) appctx.Response {
	var (
		ctx      = tracer.SpanStart(dctx.Request.Context(), "ucase.get_all_user")
		response []user.ResponseGetUser
		lf       = logger.NewFields(
			logger.EventName("GetAllUser"),
			logger.Any("querry", fmt.Sprintf("*")),
		)
	)

	authInfo := helper.AuthInfoFromContext(ctx)
	// fmt.Println("INI ROLE APA :", authInfo.Role)

	if authInfo.Role == "student" {
		logger.ErrorWithContext(ctx, fmt.Sprintf("user is not an admin"), lf...)
		return *appctx.NewResponse().WithCode(consts.CodeAuthenticationFailure).WithMessage("Unauthorized!")
	}

	request := u.connectorAdmin.GetAll()
	err := u.connectorAdmin.Send(ctx, request, &response)

	if err != nil {
		tracer.SpanError(ctx, err)
		logger.ErrorWithContext(ctx, fmt.Sprintf("all users not connector %v", err), lf...)
		return *appctx.NewResponse().WithCode(consts.CodeInternalServerError).WithMessage("Gel all users is empty!")
	}

	return *appctx.NewResponse().WithCode(consts.CodeSuccess).WithData(response).WithMessage("Get All Users Success")
}
