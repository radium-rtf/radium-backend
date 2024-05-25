package create

import "github.com/google/uuid"

type DialogueCreate struct {
	UserId uuid.UUID `json:"userId" validate:"required"`
}
