package entity

import (
	"context"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"time"
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
		m.Id = uuid.New()
	case *bun.UpdateQuery:
		m.UpdatedAt = time.Now()
	}
	return nil
}
