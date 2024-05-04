package postgres

import (
	"github.com/radium-rtf/radium-backend/internal/radium/entity"
	"github.com/uptrace/bun"
)

func initDB(db *bun.DB) {
	db.RegisterModel((*entity.CourseAuthor)(nil))
	db.RegisterModel((*entity.CourseCoauthor)(nil))
	db.RegisterModel((*entity.Student)(nil))
	db.RegisterModel((*entity.GroupCourse)(nil))
}
