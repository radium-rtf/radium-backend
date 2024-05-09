package user

import (
	"github.com/go-chi/chi/v5"
	"github.com/radium-rtf/radium-backend/internal/radium/httpserver/handlers/user/internal/search"
	mwAuth "github.com/radium-rtf/radium-backend/internal/radium/httpserver/middleware/auth"
	"github.com/radium-rtf/radium-backend/internal/radium/usecase"
)

func New(r chi.Router, useCases usecase.UseCases) {
	useCase := useCases.Account

	r.Route("/v1/user", func(r chi.Router) {
		r.Use(mwAuth.Required(useCases.Deps.TokenManager))
		r.Get("/search", search.New(useCase))
	})
}
