package user

import (
	"github.com/go-chi/chi/v5"
	"github.com/radium-rtf/radium-backend/internal/radium/httpserver/middleware/auth"
	"github.com/radium-rtf/radium-backend/internal/wave/httpserver/handlers/user/internal/get"
	"github.com/radium-rtf/radium-backend/internal/wave/usecase"
)

func New(r chi.Router, useCases usecase.UseCases) {
	useCase := useCases.User
	tokenManager := useCases.Deps.TokenManager

	r.Route("/v1/user", func(r chi.Router) {
		r.Use(auth.UserId(tokenManager))
		r.Get("/token", get.NewToken(useCase))
	})
}
