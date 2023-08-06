package entity

import "github.com/google/uuid"

type (
	Review struct {
		AnswerId   uuid.UUID `gorm:"primaryKey; type:uuid"`
		ReviewerId uuid.UUID `gorm:"type:uuid"`

		Score   float32
		Comment string
	}
)
