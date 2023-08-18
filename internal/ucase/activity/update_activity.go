package activity

import (
	"fmt"
	"strconv"

	"gitlab.com/willysihombing/task-c3/internal/helper"

	"github.com/gorilla/mux"
	"gitlab.com/willysihombing/task-c3/internal/appctx"
	"gitlab.com/willysihombing/task-c3/internal/connector/supabase/activity"
	"gitlab.com/willysihombing/task-c3/internal/consts"
	"gitlab.com/willysihombing/task-c3/internal/presentations"
	"gitlab.com/willysihombing/task-c3/internal/ucase/contract"
	"gitlab.com/willysihombing/task-c3/pkg/logger"
	"gitlab.com/willysihombing/task-c3/pkg/tracer"
)

type updateActivity struct {
	connectorIntern activity.Connectorer
}

func UpdateActivity(connectorIntern activity.Connectorer) contract.UseCase {
	return &updateActivity{
		connectorIntern: connectorIntern,
	}
}

func (u *updateActivity) Serve(dctx *appctx.Data) appctx.Response {
	var (
		ctx      = tracer.SpanStart(dctx.Request.Context(), "ucase.update_activity")
		i        = mux.Vars(dctx.Request)["id"]
		param    presentations.UpdateActivity
		response interface{}
		err      = dctx.Cast(&param)
		lf       = logger.NewFields(
			logger.EventName("CreateNewActivity"),
			logger.Any("querry", fmt.Sprintf("*")),
		)
	)

	authInfo := helper.AuthInfoFromContext(ctx)

	if authInfo.Role != "student" {
		logger.ErrorWithContext(ctx, fmt.Sprintf("user is not an students"), lf...)
		return *appctx.NewResponse().WithCode(consts.CodeAuthenticationFailure).WithMessage("Unauthorized!")
	}

	id, err := strconv.Atoi(i)
	if err != nil {
		logger.ErrorWithContext(ctx, fmt.Sprintf("update activity is not connector erorr %v", err), lf...)
		return *appctx.NewResponse().WithCode(consts.CodeInternalServerError).WithMessage("Update activity is error!")
	}

	param.UserId = authInfo.ID

	request := u.connectorIntern.UpdateByID(id, param)
	err = u.connectorIntern.Send(ctx, request, response)

	if err != nil {
		tracer.SpanError(ctx, err)
		logger.ErrorWithContext(ctx, fmt.Sprintf("update activity not connector %v", err), lf...)
		return *appctx.NewResponse().WithCode(consts.CodeInternalServerError).WithMessage("activity is not update")
	}
	return *appctx.NewResponse().WithCode(consts.CodeSuccess).WithMessage("Update activity success")
}
