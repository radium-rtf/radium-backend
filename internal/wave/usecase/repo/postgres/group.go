package postgres

import (
	"github.com/radium-rtf/radium-backend/pkg/postgres"
	"github.com/uptrace/bun"
)

type Group struct {
	db *bun.DB
}

func NewGroupRepo(pg *postgres.Postgres) Group {
	return Group{db: pg.DB}
}

func (r Group) Create() error {
	return nil
}
