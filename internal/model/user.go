package model

import (
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/entity"
)

type (
	User struct {
		Id     uuid.UUID `json:"id"`
		Email  string    `json:"email"`
		Name   string    `json:"name"`
		Avatar string    `json:"avatar"`

		Roles []Role `json:"roles"`
	}
)

func NewUser(user *entity.User) User {
	roles := make([]Role, 0, 1)
	if user.IsTeacher {
		roles = append(roles, TeacherRole)
	}
	if user.IsAuthor {
		roles = append(roles, AuthorRole)
	}

	return User{
		Id:     user.Id,
		Email:  user.Email,
		Name:   user.Name,
		Avatar: user.Avatar,
		Roles:  roles,
	}
}
