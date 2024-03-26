package postgres

import (
	"github.com/radium-rtf/radium-backend/pkg/postgres"
)

type Repositories struct {
	Channel   Channel
	Content   Content
	Dialogue  Dialogue
	GroupChat GroupChat
	Message   Message
	Post      Post
}

func NewRepositories(pg *postgres.Postgres) Repositories {
	return Repositories{}
}
