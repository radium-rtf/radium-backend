package entity

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type (
	Link struct {
		bun.BaseModel `bun:"table:links"`
		DBModel
		Name     string    `gorm:"type:string; not null"`
		Link     string    `gorm:"type:string; not null"`
		CourseId uuid.UUID `gorm:"type:uuid; not null"`
	}
)
