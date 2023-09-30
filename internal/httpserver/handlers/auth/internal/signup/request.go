package signup

import (
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/entity"
)

type SignUp struct {
	Email    string `json:"email" validate:"email"`
	Password string `json:"password" validate:"password"`
	Name     string `json:"name" validate:"required,min=1,max=30"`
}

func (s SignUp) toUser() *entity.User {
	return &entity.User{
		DBModel:  entity.DBModel{Id: uuid.New()},
		Email:    s.Email,
		Password: s.Password,
		Name:     s.Name,
	}
}
