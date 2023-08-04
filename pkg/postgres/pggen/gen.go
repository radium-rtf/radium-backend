package pggen

import (
	"github.com/radium-rtf/radium-backend/internal/entity"
	"gorm.io/gen"
)

func Gen() error {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./pkg/postgres/db",
		Mode:    gen.WithQueryInterface,
	})
	g.ApplyBasic(
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

		entity.Group{},
		entity.Teacher{},
		entity.TeacherCourse{},

		entity.AnswerReview{},
		entity.CodeReview{},
	)
	g.Execute()
	return nil
}
