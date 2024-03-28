package postgres

import (
	"github.com/radium-rtf/radium-backend/pkg/postgres"
	"github.com/uptrace/bun"
)

type Content struct {
	db *bun.DB
}

func NewContentRepo(pg *postgres.Postgres) Content {
	return Content{db: pg.DB}
}

func (r Content) Create() error {
	return nil
}
