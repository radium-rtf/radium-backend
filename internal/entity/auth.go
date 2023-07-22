package entity

import (
	"time"

	"github.com/google/uuid"
)

type (
	Session struct {
		RefreshToken string    `json:"refreshToken; not null"`
		ExpiresIn    time.Time `json:"expiresIn; not null"`
		UserId       uuid.UUID `json:"userId" gorm:"type:uuid; not null"`
	}
)
