package updaterole

import (
	"github.com/google/uuid"
	entity2 "github.com/radium-rtf/radium-backend/internal/radium/entity"
)

type Roles struct {
	IsAuthor   bool `json:"isAuthor"`
	IsTeacher  bool `json:"isTeacher"`
	IsCoauthor bool `json:"isCoauthor"`
	IsAdmin    bool `json:"isAdmin"`
}

func (r Roles) ToUser(userId uuid.UUID) *entity2.User {
	return &entity2.User{
		DBModel: entity2.DBModel{Id: userId},
		Roles: &entity2.Roles{
			IsTeacher:  r.IsTeacher,
			IsAuthor:   r.IsAuthor,
			IsCoauthor: r.IsCoauthor,
			IsAdmin:    r.IsAdmin,
		},
	}
}
