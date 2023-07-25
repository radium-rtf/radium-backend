package entity

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrPageNotFound = errors.New("страницы курса не найдены")
)

type (
	Page struct {
		DBModel
		Name     string    `gorm:"type:string; not null"`
		Slug     string    `gorm:"type:string; not null"`
		ModuleId uuid.UUID `gorm:"type:uuid; not null"`
		Order    float64   `gorm:"not null; default:0"`
		Sections []*Section
	}
)
