// Package migration
package migration

import (
	"time"

	"gitlab.com/willysihombing/task-c3/pkg/postgres"

	"gitlab.com/willysihombing/task-c3/internal/appctx"
	"gitlab.com/willysihombing/task-c3/pkg/logger"
)

func MigrateDatabase() {
	cfg, e := appctx.NewConfig()

	if e != nil {
		logger.Fatal(e)
	}

	postgres.DatabaseMigration(&postgres.Config{
		Host:         cfg.WriteDB.Host,
		Port:         cfg.WriteDB.Port,
		Name:         cfg.WriteDB.Name,
		User:         cfg.WriteDB.User,
		Password:     cfg.WriteDB.Pass,
		Timeout:      time.Duration(cfg.WriteDB.TimeoutSecond) * time.Second,
		MaxIdleConns: cfg.WriteDB.MaxIdle,
		MaxOpenConns: cfg.WriteDB.MaxOpen,
		MaxLifetime:  time.Duration(cfg.WriteDB.MaxLifeTimeMS) * time.Millisecond,
	})
}
