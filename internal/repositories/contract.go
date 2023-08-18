package repositories

import (
	"context"
	"database/sql"

	"gitlab.com/willysihombing/task-c3/internal/entity"
)

// DBTransaction contract database transaction
type DBTransaction interface {
	ExecTX(ctx context.Context, options *sql.TxOptions, fn func(context.Context, StoreTX) (int64, error)) (int64, error)
}

// StoreTX data store transaction contract
type StoreTX interface {
	// Create your function contract here
}

type Client interface {
	Get(context.Context, entity.ClientApiKey) error
}
