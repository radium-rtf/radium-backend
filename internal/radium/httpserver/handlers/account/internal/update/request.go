package update

import (
	"database/sql"

	"github.com/google/uuid"
	entity2 "github.com/radium-rtf/radium-backend/internal/radium/entity"
	"github.com/radium-rtf/radium-backend/internal/radium/model"
)

type User struct {
	Name    string         `json:"name" validate:"max=25"`
	Avatar  string         `json:"avatar" validate:"url"`
	Contact *model.Contact `json:"contact"`
}

func (u User) ToUser(userId uuid.UUID) *entity2.User {
	var contact *entity2.Contact
	if u.Contact != nil {
		contact = &entity2.Contact{
			Name:   u.Contact.Name,
			Link:   u.Contact.Link,
			UserId: userId,
			User:   &entity2.User{DBModel: entity2.DBModel{Id: userId}},
		}
	}

	return &entity2.User{
		DBModel: entity2.DBModel{Id: userId},
		Avatar:  sql.NullString{String: u.Avatar, Valid: len(u.Avatar) != 0},
		Name:    u.Name,
		Contact: contact,
	}
}
