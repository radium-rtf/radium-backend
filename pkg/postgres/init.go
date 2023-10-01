package postgres

import (
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/uptrace/bun"
)

func initDB(db *bun.DB) {
	db.RegisterModel((*entity.CourseAuthor)(nil))
	db.RegisterModel((*entity.CourseStudent)(nil))
}
