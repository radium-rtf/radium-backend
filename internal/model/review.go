package model

import (
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/entity"
)

type (
	Review struct {
		AnswerId   uuid.UUID `json:"answerId"`
		ReviewerId uuid.UUID `json:"reviewerId"`

		Score   float64 `json:"score"`
		Comment string  `json:"comment"`
	}
)

func NewReview(review *entity.Review) *Review {
	if review == nil {
		return nil
	}
	return &Review{
		AnswerId:   review.AnswerId,
		ReviewerId: review.ReviewerId,
		Score:      review.Score,
		Comment:    review.Comment,
	}
}
