package usecase

import (
	"context"
	"errors"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/internal/usecase/repo/postgres"
)

type ReviewUseCase struct {
	reviewRepo postgres.Review
	section    postgres.Section
}

func NewReviewUseCase(reviewRepo postgres.Review, sectionRepo postgres.Section) ReviewUseCase {
	return ReviewUseCase{reviewRepo: reviewRepo, section: sectionRepo}
}

func (r ReviewUseCase) Create(ctx context.Context, review *entity.Review) (*entity.Review, error) {
	section, err := r.section.GetByAnswerId(ctx, review.AnswerId)
	if err != nil {
		return nil, err
	}

	if section.MaxScore == 0 && review.Score != 0 {
		return nil, errors.New("нельзся поставить больше баллов, чем возможно")
	}

	if section.MaxScore == 0 {
		review.Score = 0
	} else {
		review.Score = review.Score / float64(section.MaxScore)
	}

	if review.Score > 1 {
		return nil, errors.New("нельзся поставить больше баллов, чем возможно")
	}

	return r.reviewRepo.Create(ctx, review)
}
