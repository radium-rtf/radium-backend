package create

import (
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/entity"
)

type Review struct {
	Id      uuid.UUID `json:"answerId"`
	Score   float32   `json:"score" validate:"numeric,min=0,max=1"`
	Comment string    `json:"comment" validate:"max=400"`
}

func (r Review) toReview(reviewerId uuid.UUID) *entity.Review {
	return &entity.Review{
		AnswerId:   r.Id,
		Score:      r.Score,
		Comment:    r.Comment,
		ReviewerId: reviewerId,
	}
}
