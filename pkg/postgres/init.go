package postgres

import (
	entity2 "github.com/radium-rtf/radium-backend/internal/radium/entity"
	"github.com/uptrace/bun"
)

func initDB(db *bun.DB) {
	db.RegisterModel((*entity2.CourseAuthor)(nil))
	db.RegisterModel((*entity2.CourseStudent)(nil))
	db.RegisterModel((*entity2.CourseCoauthor)(nil))
	db.RegisterModel((*entity2.GroupStudent)(nil))
	db.RegisterModel((*entity2.GroupCourse)(nil))
}
