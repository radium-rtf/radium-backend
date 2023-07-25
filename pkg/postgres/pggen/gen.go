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

		entity.Page{},
		entity.Course{},
		entity.Module{},
		entity.Link{},

		entity.Section{},
		entity.TextSection{},
		entity.ChoiceSection{},
		entity.MultiChoiceSection{},
		entity.ShortAnswerSection{},

		entity.Answer{},
		entity.ChoiceSectionAnswer{},
		entity.ShortAnswerSectionAnswer{},
		entity.MultichoiceSectionAnswer{},

		entity.Group{},
		entity.Teacher{},
		entity.TeacherCourse{},

		entity.AnswerReview{},
	)
	g.Execute()
	return nil
}
