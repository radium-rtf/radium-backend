package model

import (
	"github.com/google/uuid"
	"time"
)

type (
	Tokens struct {
		User         *User     `json:"user"`
		AccessToken  string    `json:"accessToken"`
		RefreshToken uuid.UUID `json:"refreshToken"`
		ExpiresIn    time.Time `json:"expiresIn"`
	}
)
