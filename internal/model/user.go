package model

import (
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/entity"
)

type (
	User struct {
		Id        uuid.UUID `json:"id"`
		Email     string    `json:"email"`
		Name      string    `json:"name"`
		Avatar    string    `json:"avatar"`
		IsTeacher bool      `json:"isTeacher"`
	}
)

func NewUser(user *entity.User) User {
	return User{
		Id:        user.Id,
		Email:     user.Email,
		Name:      user.Name,
		Avatar:    user.Avatar,
		IsTeacher: user.IsTeacher,
	}
}
