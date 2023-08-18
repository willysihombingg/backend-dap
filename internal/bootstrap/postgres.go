package bootstrap

import (
	"time"

	"gitlab.com/willysihombing/task-c3/pkg/logger"
	"gitlab.com/willysihombing/task-c3/pkg/mariadb"
	"gitlab.com/willysihombing/task-c3/pkg/postgres"

	config "gitlab.com/willysihombing/task-c3/internal/appctx"
)

func RegistryPostgres(cfg *config.Database) mariadb.Adapter {
	db, err := postgres.NewAdapter(&postgres.Config{
		Host:         cfg.Host,
		Name:         cfg.Name,
		Password:     cfg.Pass,
		Port:         cfg.Port,
		User:         cfg.User,
		Timeout:      time.Duration(cfg.TimeoutSecond) * time.Second,
		MaxOpenConns: cfg.MaxOpen,
		MaxIdleConns: cfg.MaxIdle,
		MaxLifetime:  time.Duration(cfg.MaxLifeTimeMS) * time.Millisecond,
	})

	if err != nil {
		logger.Fatal(
			err,
			logger.EventName("db"),
			logger.Any("host", cfg.Host),
			logger.Any("port", cfg.Port),
		)
	}

	return db
}
