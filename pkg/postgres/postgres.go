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

	gormDb.AutoMigrate(
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
	g.Execute()

	Q := db.Use(gormDb)

	// m := Q.Section
	// w, _ := uuid.Parse("3d1e8a0b-ecee-4270-90f0-bdc41e5ba2df")
	// t, _ := m.WithContext(context.Background()).Debug().Where(m.Id.Eq(w)).Preload(m.ChoiceSection).Preload(m.TextSection).Preload(m.MultiChoiceSection).Take()

	// // fmt.Printf("%+v", t[0])
	// b, _ := json.Marshal(t)

	// print(string(b))
	return Q, nil
}
