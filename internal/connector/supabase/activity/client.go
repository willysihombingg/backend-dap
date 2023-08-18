package activity

import "gitlab.com/willysihombing/task-c3/internal/appctx"

type client struct {
	cfg *appctx.Config
}

func NewClient(cfg *appctx.Config) Connectorer {
	return &client{
		cfg: cfg,
	}
}
