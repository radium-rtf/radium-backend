package usecase

import (
	"context"
	"errors"

	"github.com/radium-rtf/radium-backend/internal/radium/entity"
	postgres "github.com/radium-rtf/radium-backend/internal/radium/usecase/repo/postgres"
)

type ReviewUseCase struct {
	reviewRepo postgres.Review
	section    postgres.Section
	answer     postgres.Answer
}

func NewReviewUseCase(reviewRepo postgres.Review, answerRepo postgres.Answer) ReviewUseCase {
	return ReviewUseCase{reviewRepo: reviewRepo, answer: answerRepo}
}

func (r ReviewUseCase) Create(ctx context.Context, review *entity.Review) (*entity.Review, error) {
	answer, err := r.answer.GetById(ctx, review.AnswerId)
	if err != nil {
		return nil, err
	}

	section := answer.Section
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

	review.Answer = answer
	return r.reviewRepo.Create(ctx, review)
}
