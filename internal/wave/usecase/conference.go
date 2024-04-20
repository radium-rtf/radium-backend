package usecase

import (
	postgres2 "github.com/radium-rtf/radium-backend/internal/wave/usecase/repo/postgres"
)

type ConferenceUseCase struct {
	conference postgres2.Conference
}

func NewConferenceUseCase(conferenceRepo postgres2.Conference) ConferenceUseCase {
	return ConferenceUseCase{conference: conferenceRepo}
}
