package postgres

import (
	"github.com/radium-rtf/radium-backend/pkg/postgres"
)

type Repositories struct {
	Channel
	Content
	Dialogue
	GroupChat
	Message
	Post
}

func NewRepositories(pg *postgres.Postgres) Repositories {
	return Repositories{
		Channel:   NewChannelRepo(pg),
		Content:   NewContentRepo(pg),
		Dialogue:  NewDialogueRepo(pg),
		GroupChat: NewGroupChatRepo(pg),
		Message:   NewMessageRepo(pg),
		Post:      NewPostRepo(pg),
	}
}
