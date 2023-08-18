package activity

import (
	"fmt"
	"strconv"

	"gitlab.com/willysihombing/task-c3/internal/helper"

	"github.com/gorilla/mux"
	"gitlab.com/willysihombing/task-c3/internal/appctx"
	"gitlab.com/willysihombing/task-c3/internal/connector/supabase/activity"
	"gitlab.com/willysihombing/task-c3/internal/consts"
	"gitlab.com/willysihombing/task-c3/internal/ucase/contract"
	"gitlab.com/willysihombing/task-c3/pkg/logger"
	"gitlab.com/willysihombing/task-c3/pkg/tracer"
)

type deleteActivity struct {
	connectorIntern activity.Connectorer
}

func DeleteActivity(connectorIntern activity.Connectorer) contract.UseCase {
	return &deleteActivity{
		connectorIntern: connectorIntern,
	}
}

func (u *deleteActivity) Serve(dctx *appctx.Data) appctx.Response {
	var (
		ctx      = tracer.SpanStart(dctx.Request.Context(), "ucase.delete_activity")
		params   = mux.Vars(dctx.Request)["id"]
		response interface{}
		lf       = logger.NewFields(
			logger.EventName("Delete Activity"),
			logger.Any("querry", fmt.Sprintf(params)),
		)
	)
	fmt.Println(params)

	authInfo := helper.AuthInfoFromContext(ctx)

	if authInfo.Role != "student" {
		logger.ErrorWithContext(ctx, fmt.Sprintf("user is not an admin"), lf...)
		return *appctx.NewResponse().WithCode(consts.CodeAuthenticationFailure).WithMessage("Unauthorized!")
	}

	id, err := strconv.Atoi(params)
	if err != nil {
		logger.ErrorWithContext(ctx, fmt.Sprintf("activity is not deleted %v", err), lf...)
		return *appctx.NewResponse().WithCode(consts.CodeInternalServerError).WithMessage("delete activity error!")
	}
	fmt.Println("ini ID :", id)

	request := u.connectorIntern.Delete(id)
	err = u.connectorIntern.Send(ctx, request, response)

	if err != nil {
		tracer.SpanError(ctx, err)
		logger.ErrorWithContext(ctx, fmt.Sprintf("error delete is activity %v", err), lf...)
		return *appctx.NewResponse().WithCode(consts.CodeInternalServerError).WithMessage("delete activity is fail")
	}

	return *appctx.NewResponse().WithCode(consts.CodeSuccess).WithMessage("Delete is success")

}
