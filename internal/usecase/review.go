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

func (r ReviewUseCase) CreateAnswerReview(ctx context.Context, review *entity.AnswerReview) (*entity.AnswerReview, error) {
	return r.reviewRepo.CreateAnswerReview(ctx, review)
}

func (r ReviewUseCase) CreateCodeReview(ctx context.Context, review *entity.CodeReview) (*entity.CodeReview, error) {
	return r.reviewRepo.CreateCodeReview(ctx, review)
}
