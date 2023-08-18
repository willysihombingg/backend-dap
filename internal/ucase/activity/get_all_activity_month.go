package activity

import (
	"fmt"
	"time"

	"gitlab.com/willysihombing/task-c3/internal/appctx"
	"gitlab.com/willysihombing/task-c3/internal/connector/supabase/activity"
	"gitlab.com/willysihombing/task-c3/internal/consts"
	"gitlab.com/willysihombing/task-c3/internal/helper"
	"gitlab.com/willysihombing/task-c3/internal/ucase/contract"
	"gitlab.com/willysihombing/task-c3/pkg/logger"
	"gitlab.com/willysihombing/task-c3/pkg/tracer"
)

type getAllActivityMonth struct {
	connectorIntern activity.Connectorer
}

func GetAllActivityMonth(connectorIntern activity.Connectorer) contract.UseCase {
	return &getAllActivityMonth{
		connectorIntern: connectorIntern,
	}
}

func (u *getAllActivityMonth) Serve(dctx *appctx.Data) appctx.Response {
	var (
		ctx = tracer.SpanStart(dctx.Request.Context(), "ucase.documents_stamp_category")

		uid_param = dctx.Request.URL.Query().Get("user_id")
		response  []activity.ResponseGetActivity
		result    []activity.ActivityByMonth
		lf        = logger.NewFields(
			logger.EventName("GetAllActivityMonth"),
			logger.Any("querry", fmt.Sprintf("*")),
		)
	)
	fmt.Println("cek user_id : ", uid_param)

	authInfo := helper.AuthInfoFromContext(ctx)

	if !(authInfo.Role != "buddy" || authInfo.Role != "student") {
		logger.ErrorWithContext(ctx, fmt.Sprintf("user is not an students or buddy"), lf...)
		return *appctx.NewResponse().WithCode(consts.CodeAuthenticationFailure).WithMessage("Unauthorized!")
	}

	if authInfo.Role == "buddy" {
		if uid_param == "" {
			logger.ErrorWithContext(ctx, fmt.Sprintf("uid_param cannot be empty "), lf...)
			return *appctx.NewResponse().WithCode(consts.CodeAuthenticationFailure).WithMessage("Get all is empty!")
		}
	}

	request := u.connectorIntern.GetAllActivityByMonth(uid_param)
	err := u.connectorIntern.Send(ctx, request, &response)

	for _, val := range response {

		if !contains(result, val.Date) {
			result = append(result, activity.ActivityByMonth{ID: uint64(val.ID), Date: val.Date, Verified: val.Verified, Highlight: activity.Highlight{Color: chooseColor(val.Verified), Fillmode: "solid"}})
		}
	}

	if err != nil {
		tracer.SpanError(ctx, err)
		logger.ErrorWithContext(ctx, fmt.Sprintf("all users not connector %v", err), lf...)
		return *appctx.NewResponse().WithCode(consts.CodeInternalServerError).WithMessage("Gel all users is empty!")
	}

	return *appctx.NewResponse().WithCode(consts.CodeSuccess).WithData(&result).WithMessage("Get All Users Success")
}

func contains(s []activity.ActivityByMonth, date time.Time) bool {
	for _, v := range s {
		if v.Date.Day() == date.Day() {
			return true
		}
	}

	return false
}
func chooseColor(v bool) string {
	if v {
		return "green"
	} else {
		return "gray"
	}
}
