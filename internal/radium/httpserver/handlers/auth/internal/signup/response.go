package signup

import "time"

type Response struct {
	ExpiresAt time.Time `json:"expiresAt"`
	Email     string    `json:"email"`
}
