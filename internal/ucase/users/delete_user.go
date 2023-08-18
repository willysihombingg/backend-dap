package user

import (
	"fmt"
	"strconv"

	"github.com/gorilla/mux"
	"gitlab.com/willysihombing/task-c3/internal/appctx"
	"gitlab.com/willysihombing/task-c3/internal/connector/supabase/user"
	"gitlab.com/willysihombing/task-c3/internal/consts"
	"gitlab.com/willysihombing/task-c3/internal/helper"
	"gitlab.com/willysihombing/task-c3/internal/ucase/contract"
	"gitlab.com/willysihombing/task-c3/pkg/logger"
	"gitlab.com/willysihombing/task-c3/pkg/tracer"
)

type deleteUser struct {
	connectorAdmin user.Connectorer
}

func DeleteUser(connectorAdmin user.Connectorer) contract.UseCase {
	return &deleteUser{
		connectorAdmin: connectorAdmin,
	}
}

func (u *deleteUser) Serve(dctx *appctx.Data) appctx.Response {
	var (
		ctx      = tracer.SpanStart(dctx.Request.Context(), "ucase.delete_users")
		params   = mux.Vars(dctx.Request)["id"]
		response interface{}
		lf       = logger.NewFields(
			logger.EventName("Delete User"),
			logger.Any("querry", fmt.Sprintf(params)),
		)
	)

	authInfo := helper.AuthInfoFromContext(ctx)
	fmt.Println("Cek Auth :", authInfo)

	if authInfo.Role != "admin" {
		logger.ErrorWithContext(ctx, fmt.Sprintf("user is not an admin"), lf...)
		return *appctx.NewResponse().WithCode(consts.CodeAuthenticationFailure).WithMessage("Unauthorized!")
	}

	id, err := strconv.Atoi(params)
	if err != nil {
		logger.ErrorWithContext(ctx, fmt.Sprintf("users is not deleted %v", err), lf...)
		return *appctx.NewResponse().WithCode(consts.CodeInternalServerError).WithMessage("delete users error!")
	}

	request := u.connectorAdmin.Delete(id)
	err = u.connectorAdmin.Send(ctx, request, response)

	if err != nil {
		tracer.SpanError(ctx, err)
		logger.ErrorWithContext(ctx, fmt.Sprintf("error delete is users %v", err), lf...)
		return *appctx.NewResponse().WithCode(consts.CodeInternalServerError).WithMessage("delete users is fail")
	}

	return *appctx.NewResponse().WithCode(consts.CodeSuccess).WithMessage("Delete is success")

}
