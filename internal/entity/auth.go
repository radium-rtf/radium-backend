package entity

import (
	"time"

	"github.com/google/uuid"
)

type (
	SignIn struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	SignUp struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Name     string `json:"name"`
	}

	Session struct {
		RefreshToken string    `json:"refreshToken; not null"`
		ExpiresIn    time.Time `json:"expiresIn; not null"`
		UserId       uuid.UUID `json:"userId" gorm:"type:uuid; not null"`
	}

	RefreshToken struct {
		RefreshToken string `json:"refreshToken"`
	}

	VerificationCode struct {
		VerificationCode string `json:"verificationCode"`
	}

	VerificationResult struct {
		Success bool `json:"success"`
	}

	Tokens struct {
		AccessToken  string    `json:"accessToken"`
		RefreshToken string    `json:"refreshToken"`
		ExpiresIn    time.Time `json:"expiresIn"`
	}
)
