package entity

import (
	"github.com/uptrace/bun"
)

type (
	Content struct {
		bun.BaseModel `bun:"table:wave.contents"`
		DBModel
	}
)
