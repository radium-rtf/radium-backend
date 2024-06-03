package usecase

import (
	"github.com/radium-rtf/radium-backend/internal/radium/lib/auth"
	"github.com/radium-rtf/radium-backend/internal/wave/usecase/repo/postgres"
	"github.com/radium-rtf/radium-backend/pkg/centrifugo"
	"github.com/radium-rtf/radium-backend/pkg/filestorage"
)

type Dependencies struct {
	Repos        postgres.Repositories
	Storage      filestorage.Storage
	TokenManager auth.TokenManager
	Centrifugo   centrifugo.Centrifugo
}

type UseCases struct {
	Channel    ChannelUseCase
	Chat       ChatUseCase
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
		Chat:       NewChatUseCase(repos.Chat, deps.Centrifugo),
		Conference: NewConferenceUseCase(repos.Conference),
		Content:    NewContentUseCase(repos.Content),
		Dialogue:   NewDialogueUseCase(repos.Dialogue, repos.Chat),
		Group:      NewGroupUseCase(repos.Group, repos.Chat),
		Message:    NewMessageUseCase(repos.Message, repos.Content, deps.Centrifugo),
		User:       NewUserUseCase(deps.Centrifugo),
	}
}
