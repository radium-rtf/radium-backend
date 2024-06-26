package postgres

import (
	"context"
	"github.com/google/uuid"
	entity "github.com/radium-rtf/radium-backend/internal/radium/entity"
	"github.com/radium-rtf/radium-backend/internal/radium/lib/answer/verdict"
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
		onConflict := "score = excluded.score, updated_at = excluded.updated_at, comment = excluded.comment"
		_, err := tx.NewInsert().
			On("conflict (answer_id) do update").
			Set(onConflict).
			Model(review).
			Exec(ctx)
		if err != nil {
			return err
		}
		_, err = tx.NewUpdate().
			Model(&entity.Answer{}).
			Where("id = ?", review.AnswerId).Set("verdict = ?", verdict.REVIEWED).
			Exec(ctx)

		if err != nil {
			return err
		}

		notification := entity.Notification{
			DBModel:  entity.DBModel{Id: uuid.New()},
			AnswerId: review.AnswerId,
			UserId:   review.Answer.UserId,
			Type:     entity.ReviewNotification,
		}
		_, err = tx.NewInsert().Model(&notification).Exec(ctx)
		return err
	})
	return review, err
}

func (r Review) GetByAnswerId(ctx context.Context, id uuid.UUID) (*entity.Review, error) {
	var review = new(entity.Review)
	err := r.db.NewSelect().
		Model(review).
		Where("answer_id = ?", id).
		Limit(1).
		Scan(ctx)
	return review, err
}

func (r Review) Update(ctx context.Context, review *entity.Review) (*entity.Review, error) {
	_, err := r.db.NewUpdate().
		Model(review).
		Set("reviewer_id = ?, score = ?", review.ReviewerId, review.Score).
		Where("answer_id = ?", review.AnswerId).
		Exec(ctx)
	return review, err
}
