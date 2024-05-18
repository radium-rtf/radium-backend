package user

import (
	"github.com/go-chi/chi/v5"
	"github.com/radium-rtf/radium-backend/internal/wave/httpserver/handlers/user/internal/get"
	"github.com/radium-rtf/radium-backend/internal/wave/usecase"
)

func New(r chi.Router, useCases usecase.UseCases) {
	useCase := useCases.User

	r.Route("/v1/user", func(r chi.Router) {
		r.Get("/token", get.NewToken(useCase))
	})
}
