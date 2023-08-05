package update

import (
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/entity"
)

type User struct {
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
}

func (u User) ToUser(userId uuid.UUID) *entity.User {
	return &entity.User{
		DBModel: entity.DBModel{Id: userId},
		Avatar:  u.Avatar,
		Name:    u.Name,
	}
}
