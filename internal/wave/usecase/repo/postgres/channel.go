package postgres

import (
	"github.com/radium-rtf/radium-backend/pkg/postgres"
	"github.com/uptrace/bun"
)

type Channel struct {
	db *bun.DB
}

func NewChannelRepo(pg *postgres.Postgres) Channel {
	return Channel{db: pg.DB}
}

func (r Channel) Create() error {
	return nil
}
