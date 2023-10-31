package postgres

import (
	"context"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/internal/lib/answer/verdict"
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
	err := r.db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
		_, err := tx.NewInsert().
			On("conflict (answer_id) do update").
			Set("score = excluded.score, updated_at = excluded.updated_at").
			Model(review).
			Exec(ctx)
		if err != nil {
			return err
		}

		_, err = tx.NewUpdate().
			Model(&entity.Answer{}).
			Where("id = ?", review.AnswerId).Set("verdict = ?", verdict.REVIEWED).
			Exec(ctx)

		return err
	})
	return review, err
}
