package model

import (
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/entity"
)

type (
	Review struct {
		AnswerId   uuid.UUID `json:"answerId"`
		ReviewerId uuid.UUID `json:"reviewerId"`

		Score float64 `json:"score"`
	}
)

func NewReview(review *entity.Review) *Review {
	return &Review{
		AnswerId:   review.AnswerId,
		ReviewerId: review.ReviewerId,
		Score:      review.Score,
	}
}
