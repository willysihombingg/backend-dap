package user

import (
	"fmt"
	"strconv"

	"gitlab.com/willysihombing/task-c3/internal/helper"

	"github.com/gorilla/mux"
	"gitlab.com/willysihombing/task-c3/internal/appctx"
	"gitlab.com/willysihombing/task-c3/internal/connector/supabase/user"
	"gitlab.com/willysihombing/task-c3/internal/consts"
	"gitlab.com/willysihombing/task-c3/internal/ucase/contract"
	"gitlab.com/willysihombing/task-c3/pkg/logger"
	"gitlab.com/willysihombing/task-c3/pkg/tracer"
)

type getDetailUser struct {
	connectorAdmin user.Connectorer
}

func GetDetailUser(connectorAdmin user.Connectorer) contract.UseCase {
	return &getDetailUser{
		connectorAdmin: connectorAdmin,
	}
}

func (u *getDetailUser) Serve(dctx *appctx.Data) appctx.Response {
	var (
		ctx      = tracer.SpanStart(dctx.Request.Context(), "ucase.get_detail_users")
		params   = mux.Vars(dctx.Request)["id"]
		response []user.ResponseGetUser
		lf       = logger.NewFields(
			logger.EventName("GetDetailUser"),
			logger.Any("querry", fmt.Sprintf("*")),
		)
	)

	authInfo := helper.AuthInfoFromContext(ctx)

	if authInfo.Role != "admin" {
		logger.ErrorWithContext(ctx, fmt.Sprintf("user is not an admin"), lf...)
		return *appctx.NewResponse().WithCode(consts.CodeAuthenticationFailure).WithMessage("Unauthorized!")
	}

	id, err := strconv.Atoi(params)
	if err != nil {
		logger.ErrorWithContext(ctx, fmt.Sprintf("detail users is not connector erorr %v", err), lf...)
		return *appctx.NewResponse().WithCode(consts.CodeInternalServerError).WithMessage("Get detail is empty!")
	}

	request := u.connectorAdmin.GetDetailByID(id)
	err = u.connectorAdmin.Send(ctx, request, &response)
	fmt.Println(response)

	if err != nil {
		tracer.SpanError(ctx, err)
		logger.ErrorWithContext(ctx, fmt.Sprintf("all users is not connector erorr", err), lf...)
		return *appctx.NewResponse().WithCode(consts.CodeInternalServerError).WithMessage("Get All Users is empty")
	}

	return *appctx.NewResponse().WithCode(consts.CodeSuccess).WithData(response[0]).WithMessage("GetDetail Success")
}
