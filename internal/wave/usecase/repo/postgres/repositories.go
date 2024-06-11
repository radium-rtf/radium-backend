package postgres

import (
	"github.com/radium-rtf/radium-backend/pkg/postgres"
)

type Repositories struct {
	Channel    Channel
	Chat       Chat
	Conference Conference
	Content    Content
	Dialogue   Dialogue
	Group      Group
	Message    Message
}

func NewRepositories(pg *postgres.Postgres) Repositories {
	return Repositories{
		Channel:    NewChannelRepo(pg),
		Chat:       NewChatRepo(pg),
		Conference: NewConferenceRepo(pg),
		Content:    NewContentRepo(pg),
		Dialogue:   NewDialogueRepo(pg),
		Group:      NewGroupRepo(pg),
		Message:    NewMessageRepo(pg),
	}
}
