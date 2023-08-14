package auth

import (
	"github.com/go-chi/chi/v5"
	"github.com/radium-rtf/radium-backend/internal/httpserver/handlers/auth/internal/refresh"
	"github.com/radium-rtf/radium-backend/internal/httpserver/handlers/auth/internal/signin"
	"github.com/radium-rtf/radium-backend/internal/httpserver/handlers/auth/internal/signup"
	"github.com/radium-rtf/radium-backend/internal/usecase"
)

func New(r *chi.Mux, useCases usecase.UseCases) {
	useCase := useCases.Auth

	r.Route("/v1/auth", func(r chi.Router) {
		r.Post("/signin", signin.New(useCase))
		r.Post("/signup", signup.New(useCase))
		r.Post("/refresh", refresh.New(useCase))
	})
}
