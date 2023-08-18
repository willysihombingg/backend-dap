// Package handler
package handler

import (
	"context"
	"net/http"

	"gitlab.com/willysihombing/task-c3/internal/appctx"
	"gitlab.com/willysihombing/task-c3/internal/consts"
	"gitlab.com/willysihombing/task-c3/internal/ucase/contract"
	"gitlab.com/willysihombing/task-c3/pkg/msg"
)

// HttpRequest handler func wrapper
func HttpRequest(request *http.Request, svc contract.UseCase, conf *appctx.Config) appctx.Response {
	if !msg.GetAvailableLang(200, request.Header.Get(consts.HeaderLanguageKey)) {
		request.Header.Set(consts.HeaderLanguageKey, conf.App.DefaultLang)
	}

	ctx := context.WithValue(request.Context(), consts.CtxLang, request.Header.Get(consts.HeaderLanguageKey))

	req := request.WithContext(ctx)

	data := &appctx.Data{
		Request:     req,
		Config:      conf,
		ServiceType: consts.ServiceTypeHTTP,
	}

	return svc.Serve(data)
}
