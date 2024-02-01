package createlink

import (
	"github.com/google/uuid"
	entity2 "github.com/radium-rtf/radium-backend/internal/radium/entity"
)

type Link struct {
	Name string `json:"name" validate:"required,min=1,max=32"`
	Link string `json:"link" validate:"required,uri"`
}

func (r Link) toLink(courseId uuid.UUID) *entity2.Link {
	return &entity2.Link{
		DBModel:  entity2.DBModel{Id: uuid.New()},
		Link:     r.Link,
		Name:     r.Name,
		CourseId: courseId,
	}
}
