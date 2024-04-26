package create

import "github.com/google/uuid"

type DialogueCreate struct {
	UserId uuid.UUID `json:"userId" validate:"required"`
}

func (d DialogueCreate) GetId() (uuid.UUID, error) {
	return d.UserId, nil
}
