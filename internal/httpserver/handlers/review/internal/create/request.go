package create

import (
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/entity"
)

type Review struct {
	Id      uuid.UUID `json:"answer_id"`
	Score   float32   `json:"score"`
	Comment string    `json:"comment"`
}

func (r Review) toReview(reviewerId uuid.UUID) *entity.Review {
	return &entity.Review{
		AnswerId:   r.Id,
		Score:      r.Score,
		Comment:    r.Comment,
		ReviewerId: reviewerId,
	}
}
