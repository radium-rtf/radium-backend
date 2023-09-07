package model

import "time"

type (
	Tokens struct {
		User         User      `json:"user"`
		AccessToken  string    `json:"accessToken,omitempty"`
		RefreshToken string    `json:"refreshToken,omitempty"`
		ExpiresIn    time.Time `json:"expiresIn"`
	}
)
