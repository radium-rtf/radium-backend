package update

import (
	"database/sql"
	"github.com/google/uuid"
	entity2 "github.com/radium-rtf/radium-backend/internal/radium/entity"
)

type User struct {
	Name   string `json:"name" validate:"max=25"`
	Avatar string `json:"avatar" validate:"url"`
}

func (u User) ToUser(userId uuid.UUID) *entity2.User {
	return &entity2.User{
		DBModel: entity2.DBModel{Id: userId},
		Avatar:  sql.NullString{String: u.Avatar, Valid: len(u.Avatar) != 0},
		Name:    u.Name,
	}
}
