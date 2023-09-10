package app

import (
	"github.com/radium-rtf/radium-backend/config"
	"github.com/radium-rtf/radium-backend/pkg/postgres"
)

func openDB(cfg config.PG) (*postgres.Postgres, error) {
	return postgres.New(cfg.URL,
		postgres.MaxOpenConns(cfg.MaxOpenConns),
		postgres.ConnMaxIdleTime(cfg.ConnMaxIdleTime),
		postgres.MaxIdleConns(cfg.MaxIdleConns),
		postgres.ConnMaxLifetime(cfg.ConnMaxLifetime),
	)
}
