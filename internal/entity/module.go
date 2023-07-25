package entity

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrModulesNotFound = errors.New("модули курса не найдены")
)

type (
	Module struct {
		DBModel
		Slug     string    `json:"slug" gorm:"type:string; not null"`
		Name     string    `json:"name" gorm:"type:string; not null"`
		CourseId uuid.UUID `json:"courseId" gorm:"type:uuid; not null"`
		Order    float64   `gorm:"not null; default:0"`
		Pages    []*Page
	}
)
