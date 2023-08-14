package postgres

import (
	"context"

	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/pkg/postgres/db"
)

type Review struct {
	pg *db.Query
}

func NewReviewRepo(pg *db.Query) Review {
	return Review{pg: pg}
}

func (r Review) Create(ctx context.Context, review *entity.Review) (*entity.Review, error) {
	return review, r.pg.Review.WithContext(ctx).Save(review)
}
