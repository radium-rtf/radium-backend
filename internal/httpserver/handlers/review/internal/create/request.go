package create

import (
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/entity"
)

type Review struct {
	AnswerId uuid.UUID `json:"answerId"`
	Score    float64   `json:"score" validate:"numeric,min=0,max=1"`
}

func (r Review) toReview(reviewerId uuid.UUID) *entity.Review {
	return &entity.Review{
		AnswerId:   r.AnswerId,
		Score:      r.Score,
		ReviewerId: reviewerId,
	}
}
