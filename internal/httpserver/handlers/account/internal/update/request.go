package update

import (
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/entity"
)

type User struct {
	Name   string `json:"name" validate:"required,min=1,max=25"`
	Avatar string `json:"avatar" validate:"required,url"`
}

func (u User) ToUser(userId uuid.UUID) *entity.User {
	return &entity.User{
		DBModel: entity.DBModel{Id: userId},
		Avatar:  u.Avatar,
		Name:    u.Name,
	}
}
