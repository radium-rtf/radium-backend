package gen

import (
	"github.com/radium-rtf/radium-backend/internal/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func Gen(url string) error {
	gormDb, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		return err
	}
	g := gen.NewGenerator(gen.Config{
		OutPath: "./pkg/postgres/db",
		Mode:    gen.WithQueryInterface,
	})
	g.UseDB(gormDb)
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
