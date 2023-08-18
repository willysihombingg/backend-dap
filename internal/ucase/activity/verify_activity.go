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

type verifyActivity struct {
	connectorActivity activity.Connectorer
}

func NewVerifyActivity(connectorIntern activity.Connectorer) contract.UseCase {
	return &verifyActivity{
		connectorActivity: connectorIntern,
	}
}

func (u *verifyActivity) Serve(dctx *appctx.Data) appctx.Response {
	var (
		ctx      = tracer.SpanStart(dctx.Request.Context(), "ucase.document_stamp_category")
		param    presentations.VerifyActivity
		err      = dctx.Cast(&param)
		activity []activity.ResponseGetActivity
		Body     presentations.UpdateVerifyBody
		lf       = logger.NewFields(
			logger.EventName("CreateNewActivity"),
			logger.Any("query", fmt.Sprintf("*")),
		)
	)
	authInfo := helper.AuthInfoFromContext(ctx)
	fmt.Println("INI PARAM :", param)

	if authInfo.Role != "buddy" {
		logger.ErrorWithContext(ctx, fmt.Sprintf("user is not an buddy"), lf...)
		return *appctx.NewResponse().WithCode(consts.CodeAuthenticationFailure).WithMessage("Unauthorized!")
	}

	fmt.Println("ini param date gan ---> ", param.Date)

	date, _ := time.Parse("2006-01-02", param.Date)
	fmt.Println("parse date ---> ", date)

	rget := u.connectorActivity.GetAllByDate(date, param.UserId)

	err = u.connectorActivity.Send(ctx, rget, &activity)

	if err != nil {
		tracer.SpanError(ctx, err)
		logger.ErrorWithContext(ctx, fmt.Sprintf("all users not connector %v", err), lf...)
		return *appctx.NewResponse().WithCode(consts.CodeInternalServerError).WithMessage("Gel all users is empty!")
	}

	Body.Verified = true

	//update
	rupdate := u.connectorActivity.UpdateByDateAndId(date, param.UserId, Body)
	err = u.connectorActivity.Send(ctx, rupdate, &activity)

	if err != nil {
		tracer.SpanError(ctx, err)
		logger.ErrorWithContext(ctx, fmt.Sprintf("all users not connector %v", err), lf...)
		return *appctx.NewResponse().WithCode(consts.CodeInternalServerError).WithStatus(fmt.Sprintf("%s", err)).WithMessage("Gel all users is empty!")
	}

	return *appctx.NewResponse().WithCode(consts.CodeSuccess).WithMessage("Verify success")
}
