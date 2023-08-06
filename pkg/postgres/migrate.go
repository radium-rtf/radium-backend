package postgres

import (
	"github.com/radium-rtf/radium-backend/internal/entity"
	"gorm.io/gorm"
)

func migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		entity.User{},
		entity.Session{},

		entity.Course{},
		entity.Link{},
		entity.Module{},
		entity.Page{},

		entity.Section{},
		entity.TextSection{},
		entity.ChoiceSection{},
		entity.MultiChoiceSection{},
		entity.ShortAnswerSection{},
		entity.AnswerSection{},
		entity.CodeSection{},

		entity.Answer{},
		entity.ChoiceSectionAnswer{},
		entity.ShortAnswerSectionAnswer{},
		entity.MultichoiceSectionAnswer{},
		entity.AnswerSectionAnswer{},
		entity.CodeSectionAnswer{},

		entity.Teacher{},
		entity.TeacherCourse{},

		entity.Review{},
	)
}
