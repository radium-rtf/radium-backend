package usecase

import (
	"context"

	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/internal/usecase/repo/postgres"
)

type ReviewUseCase struct {
	reviewRepo postgres.Review
}

func NewReviewUseCase(reviewRepo postgres.Review) ReviewUseCase {
	return ReviewUseCase{reviewRepo: reviewRepo}
}

func (r ReviewUseCase) Create(ctx context.Context, review *entity.Review) (*entity.Review, error) {
	return r.reviewRepo.Create(ctx, review)
}
