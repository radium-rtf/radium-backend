package entity

import "github.com/google/uuid"

type (
	AnswerReview struct {
		DBModel
		OwnerId uuid.UUID `gorm:"type:uuid"`
		Score   float32
	}

	CodeReview struct {
		DBModel
		OwnerId uuid.UUID `gorm:"type:uuid"`
		Score   float32
		Comment string
	}
)
