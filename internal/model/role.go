package model

import "github.com/radium-rtf/radium-backend/internal/entity"

type (
	Roles struct {
		IsAuthor   bool `json:"isAuthor"`
		IsTeacher  bool `json:"isTeacher"`
		IsCoauthor bool `json:"isCoauthor"`
	}
)

func NewRoles(roles *entity.Roles) *Roles {
	if roles == nil {
		return nil
	}
	return &Roles{
		IsTeacher:  roles.IsTeacher,
		IsAuthor:   roles.IsAuthor,
		IsCoauthor: roles.IsCoauthor,
	}
}
