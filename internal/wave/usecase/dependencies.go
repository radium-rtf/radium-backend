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
	Channel    ChannelUseCase
	Conference ConferenceUseCase
	Content    ContentUseCase
	Dialogue   DialogueUseCase
	Group      GroupUseCase
	Message    MessageUseCase

	Deps Dependencies
}

func NewUseCases(deps Dependencies) UseCases {
	repos := deps.Repos

	return UseCases{
		Deps: deps,

		Channel:    NewChannelUseCase(repos.Channel),
		Conference: NewConferenceUseCase(repos.Conference),
		Content:    NewContentUseCase(repos.Content),
		Dialogue:   NewDialogueUseCase(repos.Dialogue),
		Group:      NewGroupUseCase(repos.Group),
		Message:    NewMessageUseCase(repos.Message),
	}
}
