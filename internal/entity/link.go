package entity

import "github.com/google/uuid"

type (
	Link struct {
		DBModel
		Name     string    `gorm:"type:string; not null"`
		Link     string    `gorm:"type:string; not null"`
		CourseId uuid.UUID `gorm:"type:uuid; not null"`
	}
)
