package postgres

import (
	"context"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/pkg/postgres"
	"github.com/uptrace/bun"
)

type Review struct {
	db *bun.DB
}

func NewReviewRepo(pg *postgres.Postgres) Review {
	return Review{db: pg.DB}
}

func (r Review) Create(ctx context.Context, review *entity.Review) (*entity.Review, error) {
	_, err := r.db.NewInsert().
		On("conflict (answer_id) do update").
		Set("score = excluded.score, updated_at = excluded.updated_at").
		Model(review).
		Exec(ctx)
	return review, err
}
