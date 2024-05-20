package usecase

import (
	"github.com/radium-rtf/radium-backend/internal/wave/lib/centrifugo"
	"github.com/radium-rtf/radium-backend/internal/wave/usecase/repo/postgres"
	"github.com/radium-rtf/radium-backend/pkg/filestorage"
)

type Dependencies struct {
	Repos      postgres.Repositories
	Storage    filestorage.Storage
	Centrifugo centrifugo.Centrifugo
}

type UseCases struct {
	Channel    ChannelUseCase
	Conference ConferenceUseCase
	Content    ContentUseCase
	Dialogue   DialogueUseCase
	Group      GroupUseCase
	Message    MessageUseCase
	User       UserUseCase

	Deps Dependencies
}

func NewUseCases(deps Dependencies) UseCases {
	repos := deps.Repos

	return UseCases{
		Deps: deps,

		Channel:    NewChannelUseCase(repos.Channel),
		Conference: NewConferenceUseCase(repos.Conference),
		Content:    NewContentUseCase(repos.Content),
		Dialogue:   NewDialogueUseCase(repos.Dialogue, deps.Centrifugo),
		Group:      NewGroupUseCase(repos.Group),
		Message:    NewMessageUseCase(repos.Message, repos.Content, deps.Centrifugo),
		User:       NewUserUseCase(deps.Centrifugo),
	}
}
