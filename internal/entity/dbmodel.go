package entity

import (
	"github.com/google/uuid"
	"time"
)

type DBModel struct {
	Id        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
