package entity

import (
	"time"
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
		RefreshToken string    `json:"refresh_token"`
		ExpiresIn    time.Time `json:"expires_in"`
		UserId       uint      `json:"user_id"`
	}

	RefreshToken struct {
		RefreshToken string `json:"refresh_token"`
	}

	Tokens struct {
		AccessToken  string    `json:"access_token"`
		RefreshToken string    `json:"refresh_token"`
		ExpiresIn    time.Time `json:"expires_in"`
	}
)
