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
		Answer   *Answer   `bun:"rel:belongs-to,join:answer_id=id"`

		ReviewerId uuid.UUID
		Reviewer   *User `bun:"rel:belongs-to,join:reviewer_id=id"`

		Score float64

		Comment string

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
