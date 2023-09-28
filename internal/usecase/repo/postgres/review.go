package postgres

import (
	"context"
	"github.com/radium-rtf/radium-backend/pkg/postgres"

	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/pkg/postgres/db"
)

type Review struct {
	pg *db.Query
}

func NewReviewRepo(pg *postgres.Postgres) Review {
	return Review{pg: pg.Q}
}

func (r Review) Create(ctx context.Context, review *entity.Review) (*entity.Review, error) {
	return review, r.pg.Review.WithContext(ctx).Save(review)
}
