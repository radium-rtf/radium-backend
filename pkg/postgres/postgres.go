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
		Mode:    gen.WithDefaultQuery,
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
	) // модельки

	gormDb.AutoMigrate(
		entity.User{},
		entity.Session{},
		entity.Course{},
		entity.Link{},
		entity.Module{},
		entity.Page{},
	)
	g.Execute()

	Q := db.Use(gormDb)

	return Q, nil
}
