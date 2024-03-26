package usecase

import (
	postgres2 "github.com/radium-rtf/radium-backend/internal/wave/usecase/repo/postgres"
)

type ChannelUseCase struct {
	channel postgres2.Channel
}

func NewChannelUseCase(channelRepo postgres2.Channel) ChannelUseCase {
	return ChannelUseCase{channel: channelRepo}
}
