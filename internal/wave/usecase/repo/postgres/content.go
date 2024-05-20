package postgres

import (
	"context"

	"github.com/radium-rtf/radium-backend/internal/wave/entity"
	"github.com/radium-rtf/radium-backend/pkg/postgres"
	"github.com/uptrace/bun"
)

type Content struct {
	db *bun.DB
}

func NewContentRepo(pg *postgres.Postgres) Content {
	return Content{db: pg.DB}
}

func (r Content) Create(ctx context.Context, content *entity.Content) error {
	return r.db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
		_, err := tx.NewInsert().Model(content).Exec(ctx)
		return err
	})
}
