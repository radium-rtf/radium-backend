package signup

import "github.com/radium-rtf/radium-backend/internal/entity"

type Request struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

func (s Request) ToUser() *entity.User {
	return &entity.User{Email: s.Email, Password: s.Password, Name: s.Name}
}
