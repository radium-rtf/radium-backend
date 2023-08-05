package create

import (
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/entity"
)

type Answer struct {
	Id    uuid.UUID `json:"id"`
	Score float32   `json:"score"`
}

type Code struct {
	Id      uuid.UUID `json:"id"`
	Score   float32   `json:"score"`
	Comment string    `json:"comment"`
}

func (r Answer) toReview() *entity.AnswerReview {
	return &entity.AnswerReview{
		OwnerId: r.Id,
		Score:   r.Score,
	}
}

func (r Code) toReview() *entity.CodeReview {
	return &entity.CodeReview{
		OwnerId: r.Id,
		Score:   r.Score,
		Comment: r.Comment,
	}
}
