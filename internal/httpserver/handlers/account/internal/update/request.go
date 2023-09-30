package update

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/entity"
)

type User struct {
	Name   string `json:"name" validate:"max=25"`
	Avatar string `json:"avatar" validate:"url"`
}

func (u User) ToUser(userId uuid.UUID) *entity.User {
	return &entity.User{
		DBModel: entity.DBModel{Id: userId},
		Avatar:  sql.NullString{String: u.Avatar, Valid: len(u.Avatar) != 0},
		Name:    u.Name,
	}
}
