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
		Slug     string    `gorm:"type:string; not null"`
		Name     string    `gorm:"type:string; not null"`
		CourseId uuid.UUID `gorm:"type:uuid; not null"`
		Order    float64   `gorm:"not null; default:0"`
		Pages    []*Page
	}
)
