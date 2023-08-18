package activity

import (
	"fmt"
	"time"

	"gitlab.com/willysihombing/task-c3/internal/helper"

	"gitlab.com/willysihombing/task-c3/internal/appctx"
	"gitlab.com/willysihombing/task-c3/internal/connector/supabase/activity"
	"gitlab.com/willysihombing/task-c3/internal/consts"
	"gitlab.com/willysihombing/task-c3/internal/presentations"
	"gitlab.com/willysihombing/task-c3/internal/ucase/contract"
	"gitlab.com/willysihombing/task-c3/pkg/logger"
	"gitlab.com/willysihombing/task-c3/pkg/tracer"
)

type createNewActivity struct {
	connectorIntern activity.Connectorer
}

func CreateActivity(connectorIntern activity.Connectorer) contract.UseCase {
	return &createNewActivity{
		connectorIntern: connectorIntern,
	}
}

func (u *createNewActivity) Serve(dctx *appctx.Data) appctx.Response {
	var (
		ctx      = tracer.SpanStart(dctx.Request.Context(), "ucase.document_stamp_category")
		param    presentations.CreateActivity
		payload  presentations.Activity
		response interface{}
		err      = dctx.Cast(&param)
		lf       = logger.NewFields(
			logger.EventName("CreateNewActivity"),
			logger.Any("query", fmt.Sprintf("*")),
		)
	)
	authInfo := helper.AuthInfoFromContext(ctx)

	if authInfo.Role != "student" {
		logger.ErrorWithContext(ctx, fmt.Sprintf("user is not an student"), lf...)
		return *appctx.NewResponse().WithCode(consts.CodeAuthenticationFailure).WithMessage("Unauthorized!")
	}

	payload.UserId = authInfo.ID
	t, _ := time.Parse("2006-01-02", param.Date)
	payload.Date = t
	payload.Activity = param.Activity
	payload.Time = param.Time
	fmt.Println(payload)

	request := u.connectorIntern.Insert(payload)
	err = u.connectorIntern.Send(ctx, request, response)

	if err != nil {
		tracer.SpanError(ctx, err)
		logger.ErrorWithContext(ctx, fmt.Sprintf("error on connector all activity %v", err), lf...)
		return *appctx.NewResponse().WithCode(consts.CodeInternalServerError).WithMessage("Create activity error!")
	}

	return *appctx.NewResponse().WithStatus("Success create Activity").WithCode(consts.CodeSuccess).WithData("Created").WithMessage("Create Activity Success")
}
