package entity

import "github.com/google/uuid"

type (
	Link struct {
		DBModel
		Name     string    `json:"name" gorm:"type:string; not null"`
		Link     string    `json:"link" gorm:"type:string; not null"`
		CourseId uuid.UUID `json:"courseId" gorm:"type:uuid; not null"`
	}

	LinkDto struct {
		Name string `json:"name"`
		Link string `json:"link"`
	}
)
