package postgres

import (
	"github.com/radium-rtf/radium-backend/pkg/postgres"
	"github.com/uptrace/bun"
)

type Post struct {
	db *bun.DB
}

func NewPostRepo(pg *postgres.Postgres) Post {
	return Post{db: pg.DB}
}

func (r Post) Create() error {
	return nil
}
