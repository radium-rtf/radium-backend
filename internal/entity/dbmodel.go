package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type DBModel struct {
	Id        uuid.UUID      `gorm:"primaryKey; default:gen_random_uuid(); type:uuid; not null" json:"id"`
	CreatedAt time.Time      `gorm:"default:now(); not null"`
	UpdatedAt time.Time      `gorm:"default:now(); not null"`
	DeletedAt gorm.DeletedAt `gorm:"index" swaggertype:"string"`
}
