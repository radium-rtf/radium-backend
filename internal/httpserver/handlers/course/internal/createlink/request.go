package createlink

import (
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/entity"
)

type Link struct {
	Name string `json:"name" validate:"required,min=1,max=15"`
	Link string `json:"link" validate:"required,url"`
}

func (r Link) toLink(courseId uuid.UUID) *entity.Link {
	return &entity.Link{
		DBModel:  entity.DBModel{Id: uuid.New()},
		Link:     r.Link,
		Name:     r.Name,
		CourseId: courseId,
	}
}
