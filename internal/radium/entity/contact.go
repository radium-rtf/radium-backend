package entity

import (
	"github.com/google/uuid"
)

type (
	Contact struct {
		Name string `validate:"required,min=1,max=64"`
		Link string `validate:"required,url"`

		User   *User `bun:"rel:belongs-to,join:user_id=id"`
		UserId uuid.UUID
	}
)
