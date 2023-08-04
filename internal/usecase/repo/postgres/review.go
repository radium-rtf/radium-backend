package postgres

import (
	"context"

	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/pkg/postgres/db"
)

type Repo struct {
	pg *db.Query
}

func NewReviewRepo(pg *db.Query) Repo {
	return Repo{pg: pg}
}

func (r Repo) CreateAnswerReview(ctx context.Context, review *entity.AnswerReview) (*entity.AnswerReview, error) {
	return review, r.pg.AnswerReview.WithContext(ctx).Create(review)
}

func (r Repo) CreateCodeReview(ctx context.Context, review *entity.CodeReview) (*entity.CodeReview, error) {
	return review, r.pg.CodeReview.WithContext(ctx).Create(review)
}
