package signup

import (
	"github.com/radium-rtf/radium-backend/internal/entity"
)

type SignUp struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

func (s SignUp) toUser() *entity.User {
	return &entity.User{Email: s.Email, Password: s.Password, Name: s.Name}
}
