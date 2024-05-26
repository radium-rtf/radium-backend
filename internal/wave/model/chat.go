package model

import (
	"github.com/google/uuid"
)

type (
	Chat struct {
		Id   uuid.UUID `json:"id"`
		Name string    `json:"name"`
		Type string    `json:"type"`
	}
)
