package entity

import (
	"context"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"time"
)

type (
	Review struct {
		bun.BaseModel `bun:"table:reviews"`

		AnswerId uuid.UUID `bun:",pk"`

		ReviewerId uuid.UUID
		Score      float64

		UpdatedAt time.Time `bun:",nullzero"`
		CreatedAt time.Time `bun:",nullzero"`
		DeletedAt time.Time `bun:",soft_delete,nullzero"`
	}

	ReviewComment struct { // несколько преподавателей оставили комментарий
		ReviewId      uuid.UUID
		CommentatorId uuid.UUID

		Comment string
	}
)

func (m *Review) BeforeAppendModel(_ context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		m.CreatedAt = time.Now()
		m.UpdatedAt = time.Now()
	case *bun.UpdateQuery:
		m.UpdatedAt = time.Now()
	}
	return nil
}
