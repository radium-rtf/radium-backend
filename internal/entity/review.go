package entity

import "github.com/google/uuid"

type (
	AnswerReview struct {
		DBModel
		OwnerId uuid.UUID `gorm:"type:uuid"`
		Score   float32
	}
)
