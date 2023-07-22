package update

import (
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/entity"
)

type Request struct {
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
}

func (r Request) ToUser(userId uuid.UUID) *entity.User {
	return &entity.User{
		DBModel: entity.DBModel{Id: userId},
		Avatar:  r.Avatar,
		Name:    r.Name,
	}
}
