package helper

import (
	"context"

	"gitlab.com/willysihombing/task-c3/internal/connector/supabase/user"
	"gitlab.com/willysihombing/task-c3/internal/consts"
)

// AuthInfoFromContext extract auth information from context
func AuthInfoFromContext(ctx context.Context) user.ResponseGetUser {
	v, ok := ctx.Value(consts.CtxUserInfo).(user.ResponseGetUser)
	if !ok {
		return user.ResponseGetUser{}
	}

	return v
}
