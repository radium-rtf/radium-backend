package chats

import (
	"github.com/go-chi/chi/v5"
	"github.com/radium-rtf/radium-backend/internal/radium/httpserver/middleware/auth"
	"github.com/radium-rtf/radium-backend/internal/wave/httpserver/handlers/chats/internal/get"
	"github.com/radium-rtf/radium-backend/internal/wave/usecase"
)

func New(r chi.Router, useCases usecase.UseCases) {
	useCase := useCases.Dialogue
	messageUseCase := useCases.Message
	tokenManager := useCases.Deps.TokenManager

	r.Route("/v1/chats", func(r chi.Router) {
		r.Use(auth.Required(tokenManager))
		r.Get("/", get.New(useCase, messageUseCase))
		r.Get("/token/{chatId}", get.NewToken(useCase))
	})
}
