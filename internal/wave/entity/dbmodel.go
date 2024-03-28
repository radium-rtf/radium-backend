package entity

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type DBModel struct {
	Id        uuid.UUID `bun:",pk"`
	UpdatedAt time.Time `bun:",nullzero"`
	CreatedAt time.Time `bun:",nullzero"`
	DeletedAt time.Time `bun:",soft_delete,nullzero"`
}

func (m *DBModel) BeforeAppendModel(_ context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		m.CreatedAt = time.Now()
	case *bun.UpdateQuery:
		m.UpdatedAt = time.Now()
	}
	return nil
}
