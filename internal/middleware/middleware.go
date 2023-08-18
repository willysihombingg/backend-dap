// Package middleware
package middleware

import (
	"net/http"

	"gitlab.com/willysihombing/task-c3/internal/appctx"
	"gitlab.com/willysihombing/task-c3/internal/consts"
)

// MiddlewareFunc is contract for middleware and must implement this type for http if need middleware http request
type MiddlewareFunc func(r *http.Request, conf *appctx.Config) int

// FilterFunc is a iterator resolver in each middleware registered
func FilterFunc(conf *appctx.Config, r *http.Request, mfs []MiddlewareFunc) int {
	for _, mf := range mfs {
		if status := mf(r, conf); status != consts.CodeSuccess {
			return status
		}
	}

	return consts.CodeSuccess
}
