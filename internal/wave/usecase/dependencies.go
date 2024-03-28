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
	Channel   ChannelUseCase
	Content   ContentUseCase
	Dialogue  DialogueUseCase
	GroupChat GroupChatUseCase
	Message   MessageUseCase
	Post      PostUseCase

	Deps Dependencies
}

func NewUseCases(deps Dependencies) UseCases {
	repos := deps.Repos

	return UseCases{
		Deps: deps,

		Channel:   NewChannelUseCase(repos.Channel),
		Content:   NewContentUseCase(repos.Content),
		Dialogue:  NewDialogueUseCase(repos.Dialogue),
		GroupChat: NewGroupChatUseCase(repos.GroupChat),
		Message:   NewMessageUseCase(repos.Message),
		Post:      NewPostUseCase(repos.Post),
	}
}
