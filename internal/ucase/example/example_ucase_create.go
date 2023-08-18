// Package example
package example

import (
	"github.com/thedevsaddam/govalidator"

	"gitlab.com/willysihombing/task-c3/internal/appctx"
	"gitlab.com/willysihombing/task-c3/internal/consts"
	"gitlab.com/willysihombing/task-c3/internal/entity"
	"gitlab.com/willysihombing/task-c3/internal/presentations"
	"gitlab.com/willysihombing/task-c3/internal/repositories"
	"gitlab.com/willysihombing/task-c3/internal/ucase/contract"
	"gitlab.com/willysihombing/task-c3/pkg/logger"
	"gitlab.com/willysihombing/task-c3/pkg/util"
)

type exampleCreate struct {
	repo repositories.Example
}

// NewPartnerCreate initialize partner cerator
func NewPartnerCreate(repo repositories.Example) contract.UseCase {
	return &exampleCreate{repo: repo}
}

// Serve partner list data
func (u *exampleCreate) Serve(data *appctx.Data) appctx.Response {

	req := presentations.ExampleCreate{}

	err := data.Cast(&req)

	if err != nil {
		logger.Error(logger.MessageFormat("[example-create] parsing body request error: %v", err))
		return *appctx.NewResponse().WithCode(consts.CodeBadRequest).WithError(err.Error())
	}

	fl := []logger.Field{
		logger.Any("request", req),
	}

	rules := govalidator.MapData{
		"name":    []string{"required", "between:3,50"},
		"email":   []string{"required", "email"},
		"address": []string{"required"},
		"phone":   []string{"phone_number"},
	}

	opts := govalidator.Options{
		Data:  &req,  // request object
		Rules: rules, // rules map
	}

	v := govalidator.New(opts)
	ev := v.ValidateStruct()

	if len(ev) != 0 {
		logger.Warn(
			logger.MessageFormat("[example-create] validate request param err: %s", util.DumpToString(ev)),
			fl...)

		return *appctx.NewResponse().WithCode(consts.CodeBadRequest).WithError(err.Error())
	}

	_, err = u.repo.Upsert(data.Request.Context(), entity.Example{
		Name:    req.Name,
		Address: req.Address,
		Email:   req.Email,
		Phone:   req.Phone,
	})

	if err != nil {
		logger.Error(logger.MessageFormat("[example-create] %v", err))

		return *appctx.NewResponse().WithCode(consts.CodeInternalServerError)
	}

	return *appctx.NewResponse().WithCode(consts.CodeSuccess)

}
