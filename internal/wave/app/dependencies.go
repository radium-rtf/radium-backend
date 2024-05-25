package app

import (
	"github.com/radium-rtf/radium-backend/config"
	"github.com/radium-rtf/radium-backend/internal/radium/lib/auth"
	"github.com/radium-rtf/radium-backend/internal/wave/usecase"
	pg "github.com/radium-rtf/radium-backend/internal/wave/usecase/repo/postgres"
	"github.com/radium-rtf/radium-backend/pkg/centrifugo"
	"github.com/radium-rtf/radium-backend/pkg/filestorage"
	"github.com/radium-rtf/radium-backend/pkg/postgres"
)

func newDependencies(storage filestorage.Storage, cfg *config.Config, db *postgres.Postgres) usecase.Dependencies {
	repositories := pg.NewRepositories(db)
	tokenManager := auth.NewManager(cfg.JWT.SigningKey)
	centrifugo := centrifugo.New(tokenManager)
	return usecase.Dependencies{
		Repos:      repositories,
		Storage:    storage,
		Centrifugo: centrifugo,
	}
}
