package repo

import (
	"context"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/pkg/postgres/db"
)

type ReviewRepo struct {
	pg *db.Query
}

func NewReviewRepo(pg *db.Query) ReviewRepo {
	return ReviewRepo{pg: pg}
}

func (r ReviewRepo) Create(ctx context.Context, review *entity.AnswerReview) (*entity.AnswerReview, error) {
	return review, r.pg.AnswerReview.WithContext(ctx).Create(review)
}
