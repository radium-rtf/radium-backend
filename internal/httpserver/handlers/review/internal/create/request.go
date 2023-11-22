package create

import (
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/entity"
)

type Review struct {
	AnswerId uuid.UUID `json:"answerId"`
	Score    uint      `json:"score" validate:"numeric,min=0"`
	Comment  string    `json:"comment" validate:"max=500"`
}

func (r Review) toReview(reviewerId uuid.UUID) *entity.Review {
	return &entity.Review{
		AnswerId:   r.AnswerId,
		Score:      float64(r.Score),
		ReviewerId: reviewerId,
		Comment:    r.Comment,
	}
}
