package model

import (
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/radium/entity"
	"time"
)

type (
	Review struct {
		AnswerId uuid.UUID `json:"answerId"`
		Reviewer *User     `json:"reviewer"`

		Score     float64   `json:"score"`
		Comment   string    `json:"comment"`
		CreatedAt time.Time `json:"createdAt"`
	}
)

func NewReview(review *entity.Review) *Review {
	if review == nil {
		return nil
	}
	return &Review{
		Reviewer:  NewUser(review.Reviewer),
		AnswerId:  review.AnswerId,
		Score:     review.Score,
		Comment:   review.Comment,
		CreatedAt: review.CreatedAt,
	}
}
