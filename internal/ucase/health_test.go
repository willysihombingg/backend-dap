// Package ucase
package ucase

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/willysihombing/task-c3/internal/appctx"
	"gitlab.com/willysihombing/task-c3/internal/consts"
)

func TestHealthCheck_Serve(t *testing.T) {
	svc := NewHealthCheck()

	t.Run("test health check", func(t *testing.T) {
		result := svc.Serve(&appctx.Data{})

		assert.Equal(t, appctx.Response{
			Code:    consts.CodeSuccess,
			Message: "ok",
		}, result)
	})
}
