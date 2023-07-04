package entity

import "github.com/google/uuid"

type Destroy struct {
	Id     uuid.UUID `json:"id"`
	IsSoft bool      `json:"isSoft"`
}
