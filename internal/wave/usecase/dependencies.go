package usecase

import (
	"github.com/radium-rtf/radium-backend/internal/wave/usecase/repo/postgres"
	"github.com/radium-rtf/radium-backend/pkg/filestorage"
)

type Dependencies struct {
	Repos   postgres.Repositories
	Storage filestorage.Storage
}

type UseCases struct {
	Deps Dependencies
}

func NewUseCases(deps Dependencies) UseCases {
	return UseCases{
		Deps: deps,
	}
}
