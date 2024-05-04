package app

import (
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/config"
	"github.com/radium-rtf/radium-backend/pkg/postgres"
)

func openDB(cfg config.PG, defaultGroupId uuid.UUID) (*postgres.Postgres, error) {
	return postgres.New(cfg.URL,
		defaultGroupId,
		postgres.MaxOpenConns(cfg.MaxOpenConns),
		postgres.ConnMaxIdleTime(cfg.ConnMaxIdleTime),
		postgres.MaxIdleConns(cfg.MaxIdleConns),
		postgres.ConnMaxLifetime(cfg.ConnMaxLifetime),
	)
}
