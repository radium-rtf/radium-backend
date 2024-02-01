package postgres

import (
	"github.com/radium-rtf/radium-backend/pkg/postgres"
)

type Repositories struct {
}

func NewRepositories(pg *postgres.Postgres) Repositories {
	return Repositories{}
}
