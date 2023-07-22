package create

import (
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/entity"
)

type Request struct {
	Id    uuid.UUID `json:"id"`
	Score float32   `json:"score"`
}

func (r Request) ToReview() *entity.AnswerReview {
	return &entity.AnswerReview{
		OwnerId: r.Id,
		Score:   r.Score,
	}
}
