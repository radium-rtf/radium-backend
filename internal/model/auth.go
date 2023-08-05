package model

import "time"

type (
	Tokens struct {
		AccessToken  string    `json:"accessToken"`
		RefreshToken string    `json:"refreshToken"`
		ExpiresIn    time.Time `json:"expiresIn"`
	}
)
