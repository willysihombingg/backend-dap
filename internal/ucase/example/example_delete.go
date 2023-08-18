// Package example
package example

import (
	"github.com/gorilla/mux"
	"github.com/spf13/cast"

	"gitlab.com/willysihombing/task-c3/internal/appctx"
	"gitlab.com/willysihombing/task-c3/internal/consts"
	"gitlab.com/willysihombing/task-c3/internal/repositories"
	"gitlab.com/willysihombing/task-c3/internal/ucase/contract"

	"gitlab.com/willysihombing/task-c3/pkg/logger"
)

type exampleDelete struct {
	repo repositories.Example
}

func NewExampleDelete(repo repositories.Example) contract.UseCase {
	return &exampleDelete{repo: repo}
}

// Serve partner list data
func (u *exampleDelete) Serve(data *appctx.Data) appctx.Response {

	id := mux.Vars(data.Request)["id"]

	err := u.repo.Delete(data.Request.Context(), cast.ToUint64(id))

	if err != nil {
		logger.Error(logger.MessageFormat("[example-delete] %v", err))

		return *appctx.NewResponse().WithCode(consts.CodeInternalServerError)
	}

	return *appctx.NewResponse().WithCode(consts.CodeSuccess)
}
