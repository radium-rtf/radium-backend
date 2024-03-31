package usecase

import (
	"context"
	"errors"

	"github.com/radium-rtf/radium-backend/internal/radium/entity"
	postgres2 "github.com/radium-rtf/radium-backend/internal/radium/usecase/repo/postgres"
)

type ReviewUseCase struct {
	reviewRepo postgres2.Review
	section    postgres2.Section
}

func NewReviewUseCase(reviewRepo postgres2.Review, sectionRepo postgres2.Section) ReviewUseCase {
	return ReviewUseCase{reviewRepo: reviewRepo, section: sectionRepo}
}

func (r ReviewUseCase) Create(ctx context.Context, review *entity.Review) (*entity.Review, error) {
	section, err := r.section.GetByAnswerId(ctx, review.AnswerId)
	if err != nil {
		return nil, err
	}

	if section.MaxScore == 0 && review.Score != 0 {
		return nil, errors.New("нельзя поставить больше баллов, чем возможно")
	}

	if section.MaxScore == 0 {
		review.Score = 0
	} else {
		review.Score = review.Score / float64(section.MaxScore)
	}

	if review.Score > 1 {
		return nil, errors.New("нельзя поставить больше баллов, чем возможно")
	}

	return r.reviewRepo.Create(ctx, review)
}
