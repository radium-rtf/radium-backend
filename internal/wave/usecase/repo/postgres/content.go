package postgres

import (
	"context"

	"github.com/google/uuid"
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

func (r Content) Update(ctx context.Context, content *entity.Content) error {
	return r.db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
		_, err := tx.NewUpdate().Model(content).
			Where("id = ?", content.Id).
			Exec(ctx)
		return err
	})
}

func (r Content) Delete(ctx context.Context, contentId uuid.UUID) error {
	return r.db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
		_, err := tx.NewDelete().Model(&entity.Content{DBModel: entity.DBModel{Id: contentId}}).
			Where("id = ?", contentId).
			Exec(ctx)
		return err
	})
}
