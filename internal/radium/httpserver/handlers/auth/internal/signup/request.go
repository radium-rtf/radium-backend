package signup

import (
	"github.com/google/uuid"
	entity2 "github.com/radium-rtf/radium-backend/internal/radium/entity"
	"strings"
)

type SignUp struct {
	Email    string `json:"email" validate:"email"`
	Password string `json:"password" validate:"password"`
	Name     string `json:"name" validate:"required,min=1,max=48"`
}

func (s SignUp) toUser() *entity2.User {
	return &entity2.User{
		DBModel:  entity2.DBModel{Id: uuid.New()},
		Email:    strings.ToLower(s.Email),
		Password: s.Password,
		Name:     s.Name,
	}
}
