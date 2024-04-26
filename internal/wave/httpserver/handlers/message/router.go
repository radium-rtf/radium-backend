package message

import (
	"github.com/go-chi/chi/v5"
	"github.com/radium-rtf/radium-backend/internal/wave/httpserver/handlers/message/internal/get"
	"github.com/radium-rtf/radium-backend/internal/wave/httpserver/handlers/message/internal/send"
	"github.com/radium-rtf/radium-backend/internal/wave/usecase"
)

func New(r chi.Router, useCases usecase.UseCases) {
	useCase := useCases.Message

	r.Route("/v1/message", func(r chi.Router) {
		r.Post("/", send.New(useCase))
	})
	r.Route("/v1/messages", func(r chi.Router) {
		r.Get("/{chatId}", get.New(useCase))
	})
}
