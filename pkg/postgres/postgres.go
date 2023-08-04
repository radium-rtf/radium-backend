package postgres

import (
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/pkg/postgres/db"
	"github.com/radium-rtf/radium-backend/pkg/postgres/pggen"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New(url string) (*db.Query, error) {
	gormDb, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = pggen.Gen()
	if err != nil {
		return nil, err
	}
	Q := db.Use(gormDb)

	err = gormDb.AutoMigrate(
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

		entity.AnswerReview{},
		entity.CodeReview{},
	)

	if err != nil {
		return nil, err
	}

	return Q, err
}
