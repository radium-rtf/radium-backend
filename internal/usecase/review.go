package usecase

import (
	"context"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/internal/usecase/repo/postgres"
)

type ReviewUseCase struct {
	reviewRepo postgres.Repo
}

func NewReviewUseCase(reviewRepo postgres.Repo) ReviewUseCase {
	return ReviewUseCase{reviewRepo: reviewRepo}
}

func (r ReviewUseCase) Create(ctx context.Context, review *entity.AnswerReview) (*entity.AnswerReview, error) {
	return r.reviewRepo.Create(ctx, review)
}
