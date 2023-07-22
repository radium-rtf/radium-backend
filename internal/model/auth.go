package model

import "time"

type (
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
