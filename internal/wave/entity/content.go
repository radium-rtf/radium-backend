package entity

import (
	"github.com/google/uuid"
	radium "github.com/radium-rtf/radium-backend/internal/radium/entity"
	"github.com/uptrace/bun"
)

type (
	Content struct {
		bun.BaseModel `bun:"table:wave.contents"`
		DBModel

		FileId uuid.UUID
		File   *radium.File `bun:"rel:belongs-to,join:file_id=id"`

		Text string
		Type string
	}
)
