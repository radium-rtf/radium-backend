package create

import (
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/entity"
)

type AnswerRequest struct {
	Id    uuid.UUID `json:"id"`
	Score float32   `json:"score"`
}

type CodeRequest struct {
	Id      uuid.UUID `json:"id"`
	Score   float32   `json:"score"`
	Comment string    `json:"comment"`
}

func (r AnswerRequest) ToReview() *entity.AnswerReview {
	return &entity.AnswerReview{
		OwnerId: r.Id,
		Score:   r.Score,
	}
}

func (r CodeRequest) ToReview() *entity.CodeReview {
	return &entity.CodeReview{
		OwnerId: r.Id,
		Score:   r.Score,
		Comment: r.Comment,
	}
}
