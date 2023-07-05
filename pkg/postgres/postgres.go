package postgres

import (
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/pkg/postgres/db"
	"gorm.io/driver/postgres"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func New(url string) (*db.Query, error) {
	gormDb, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = Gen(url)
	if err != nil {
		return nil, err
	}
	Q := db.Use(gormDb)

	return Q, err
}

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
		entity.Group{},
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
	)

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
		entity.Answer{},
		entity.ChoiceSectionAnswer{},
		entity.ShortAnswerSectionAnswer{},
		entity.MultichoiceSectionAnswer{},
	)
	if err != nil {
		return err
	}
	g.Execute()
	return nil
}
