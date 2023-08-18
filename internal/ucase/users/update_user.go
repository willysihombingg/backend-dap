package user

import (
	"fmt"
	"strconv"

	"gitlab.com/willysihombing/task-c3/internal/connector/supabase/user"
	"gitlab.com/willysihombing/task-c3/internal/helper"

	"github.com/gorilla/mux"
	"gitlab.com/willysihombing/task-c3/internal/appctx"
	"gitlab.com/willysihombing/task-c3/internal/consts"
	"gitlab.com/willysihombing/task-c3/internal/presentations"
	"gitlab.com/willysihombing/task-c3/internal/ucase/contract"
	"gitlab.com/willysihombing/task-c3/pkg/logger"
	"gitlab.com/willysihombing/task-c3/pkg/tracer"
)

type updateUser struct {
	connectorAdmin user.Connectorer
}

func UpdateUser(connectorAdmin user.Connectorer) contract.UseCase {
	return &updateUser{
		connectorAdmin: connectorAdmin,
	}
}

func (u *updateUser) Serve(dctx *appctx.Data) appctx.Response {
	var (
		ctx      = tracer.SpanStart(dctx.Request.Context(), "ucase.update_users")
		i        = mux.Vars(dctx.Request)["id"]
		param    presentations.UpdateUser
		response interface{}
		err      = dctx.Cast(&param)
		lf       = logger.NewFields(
			logger.EventName("CreateNewUser"),
			logger.Any("querry", fmt.Sprintf("*")),
		)
	)

	authInfo := helper.AuthInfoFromContext(ctx)

	if authInfo.Role != "admin" {
		logger.ErrorWithContext(ctx, fmt.Sprintf("user is not an admin"), lf...)
		return *appctx.NewResponse().WithCode(consts.CodeAuthenticationFailure).WithMessage("Unauthorized!")
	}

	id, err := strconv.Atoi(i)
	if err != nil {
		logger.ErrorWithContext(ctx, fmt.Sprintf("update users is not connector erorr %v", err), lf...)
		return *appctx.NewResponse().WithCode(consts.CodeInternalServerError).WithMessage("Update users is error!")
	}

	request := u.connectorAdmin.UpdateByID(id, param)
	err = u.connectorAdmin.Send(ctx, request, response)

	if err != nil {
		tracer.SpanError(ctx, err)
		logger.ErrorWithContext(ctx, fmt.Sprintf("update users not connector %v", err), lf...)
		return *appctx.NewResponse().WithCode(consts.CodeInternalServerError).WithMessage("Users is not update")
	}
	return *appctx.NewResponse().WithCode(consts.CodeSuccess).WithMessage("Update users success")
}
