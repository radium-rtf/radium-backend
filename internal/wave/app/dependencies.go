package app

import (
	"github.com/radium-rtf/radium-backend/config"
	"github.com/radium-rtf/radium-backend/internal/wave/lib/centrifugo"
	"github.com/radium-rtf/radium-backend/internal/wave/usecase"
	pg "github.com/radium-rtf/radium-backend/internal/wave/usecase/repo/postgres"
	"github.com/radium-rtf/radium-backend/pkg/filestorage"
	"github.com/radium-rtf/radium-backend/pkg/postgres"
)

func newDependencies(storage filestorage.Storage, cfg *config.Config, db *postgres.Postgres) usecase.Dependencies {
	repositories := pg.NewRepositories(db)
	centrifugo := centrifugo.New(cfg.Centrifugo)
	return usecase.Dependencies{
		Repos:      repositories,
		Storage:    storage,
		Centrifugo: centrifugo,
	}
}
