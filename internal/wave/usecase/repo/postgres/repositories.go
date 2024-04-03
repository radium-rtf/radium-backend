package postgres

import (
	"github.com/radium-rtf/radium-backend/pkg/postgres"
)

type Repositories struct {
	Channel  Channel
	Content  Content
	Dialogue Dialogue
	Group    Group
	Message  Message
	Post     Post
}

func NewRepositories(pg *postgres.Postgres) Repositories {
	return Repositories{
		Channel:  NewChannelRepo(pg),
		Content:  NewContentRepo(pg),
		Dialogue: NewDialogueRepo(pg),
		Group:    NewGroupRepo(pg),
		Message:  NewMessageRepo(pg),
		Post:     NewPostRepo(pg),
	}
}
