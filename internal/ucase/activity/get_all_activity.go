package activity

import (
	"fmt"
	"strconv"
	"time"

	"gitlab.com/willysihombing/task-c3/internal/helper"

	"gitlab.com/willysihombing/task-c3/internal/appctx"
	"gitlab.com/willysihombing/task-c3/internal/connector/supabase/activity"
	"gitlab.com/willysihombing/task-c3/internal/consts"
	"gitlab.com/willysihombing/task-c3/internal/ucase/contract"
	"gitlab.com/willysihombing/task-c3/pkg/logger"
	"gitlab.com/willysihombing/task-c3/pkg/tracer"
)

type getAllActivity struct {
	connectorIntern activity.Connectorer
}

func GetAllActivity(connectorIntern activity.Connectorer) contract.UseCase {
	return &getAllActivity{
		connectorIntern: connectorIntern,
	}
}

func (u *getAllActivity) Serve(dctx *appctx.Data) appctx.Response {
	var (
		ctx        = tracer.SpanStart(dctx.Request.Context(), "ucase.documents_stamp_category")
		user_id    int
		date_param = dctx.Request.URL.Query().Get("date")
		uid_param  = dctx.Request.URL.Query().Get("user_id")
		response   []activity.ResponseGetActivity
		lf         = logger.NewFields(
			logger.EventName("GetAllActivity"),
			logger.Any("querry", fmt.Sprintf("*")),
		)
	)

	authInfo := helper.AuthInfoFromContext(ctx)

	if !(authInfo.Role != "buddy" || authInfo.Role != "student") {
		logger.ErrorWithContext(ctx, fmt.Sprintf("user is not an students or buddy"), lf...)
		return *appctx.NewResponse().WithCode(consts.CodeAuthenticationFailure).WithMessage("Unauthorized!")
	}

	r, error := strconv.Atoi(uid_param)
	if error != nil {
		logger.ErrorWithContext(ctx, fmt.Sprintf("detail users is not connector erorr %v", error), lf...)
		return *appctx.NewResponse().WithCode(consts.CodeBadRequest).WithMessage("Get all is empty!")
	}
	user_id = r

	date, _ := time.Parse("2006-01-02", date_param)

	request := u.connectorIntern.GetAllByDate(date, user_id)
	err := u.connectorIntern.Send(ctx, request, &response)

	if err != nil {
		tracer.SpanError(ctx, err)
		logger.ErrorWithContext(ctx, fmt.Sprintf("all users not connector %v", err), lf...)
		return *appctx.NewResponse().WithCode(consts.CodeInternalServerError).WithMessage("Gel all users is empty!")
	}

	return *appctx.NewResponse().WithCode(consts.CodeSuccess).WithData(response).WithMessage("Get All Users Success")
}
