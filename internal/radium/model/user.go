package model

import (
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/radium/entity"
)

type (
	User struct {
		Id     uuid.UUID `json:"id"`
		Email  string    `json:"email"`
		Name   string    `json:"name"`
		Avatar string    `json:"avatar"`

		Roles *Roles `json:"roles"`
	}
)

func NewUser(user *entity.User) *User {
	if user == nil {
		return nil
	}
	return &User{
		Id:     user.Id,
		Email:  user.Email,
		Name:   user.Name,
		Avatar: user.Avatar.String,
		Roles:  NewRoles(user.Roles),
	}
}

func NewUsers(users []*entity.User) []*User {
	res := make([]*User, 0, len(users))
	for _, user := range users {

		res = append(res, NewUser(user))
	}
	return res
}
