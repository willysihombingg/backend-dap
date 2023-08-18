// Package router
package server

import (
	"context"
	"net/http"

	"gitlab.com/willysihombing/task-c3/internal/appctx"
	ucase "gitlab.com/willysihombing/task-c3/internal/ucase/contract"
)

// httpHandlerFunc abstraction for http handler
type httpHandlerFunc func(request *http.Request, svc ucase.UseCase, conf *appctx.Config) appctx.Response

// Server contract
type Server interface {
	Run(context.Context) error
	Done()
	Config() *appctx.Config
}
