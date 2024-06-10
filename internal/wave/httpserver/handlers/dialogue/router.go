package dialogue

import (
	"github.com/go-chi/chi/v5"
	"github.com/radium-rtf/radium-backend/internal/radium/httpserver/middleware/auth"
	"github.com/radium-rtf/radium-backend/internal/wave/httpserver/handlers/dialogue/internal/create"
	"github.com/radium-rtf/radium-backend/internal/wave/httpserver/handlers/dialogue/internal/get"
	"github.com/radium-rtf/radium-backend/internal/wave/usecase"
)

func New(r chi.Router, useCases usecase.UseCases) {
	useCase := useCases.Dialogue
	tokenManager := useCases.Deps.TokenManager

	r.Route("/v1/dialogue", func(r chi.Router) {
		r.Use(auth.Required(tokenManager))
		r.Post("/create", create.NewDialogue(useCase))
		r.Get("/{chatId}", get.New(useCase))
	})
}
