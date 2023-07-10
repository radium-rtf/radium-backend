package usecase

import (
	"context"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/internal/usecase/repo"
	"github.com/radium-rtf/radium-backend/pkg/postgres/db"
)

type ReviewUseCase struct {
	reviewRepo repo.ReviewRepo
}

func NewReviewUseCase(pg *db.Query) ReviewUseCase {
	return ReviewUseCase{reviewRepo: repo.NewReviewRepo(pg)}
}

func (r ReviewUseCase) Create(ctx context.Context, review *entity.AnswerReview) (*entity.AnswerReview, error) {
	return r.reviewRepo.Create(ctx, review)
}
