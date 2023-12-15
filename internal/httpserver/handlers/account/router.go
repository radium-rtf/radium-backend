package account

import (
	"github.com/go-chi/chi/v5"
	"github.com/radium-rtf/radium-backend/internal/httpserver/handlers/account/internal/courses"
	"github.com/radium-rtf/radium-backend/internal/httpserver/handlers/account/internal/get"
	"github.com/radium-rtf/radium-backend/internal/httpserver/handlers/account/internal/update"
	"github.com/radium-rtf/radium-backend/internal/httpserver/handlers/account/internal/updatepass"
	mwAuth "github.com/radium-rtf/radium-backend/internal/httpserver/middleware/auth"
	"github.com/radium-rtf/radium-backend/internal/httpserver/middleware/user"
	"github.com/radium-rtf/radium-backend/internal/usecase"
)

func New(r chi.Router, useCases usecase.UseCases) {
	useCase := useCases.Account

	r.Route("/v1/account", func(r chi.Router) {
		r.Use(mwAuth.Required(useCases.Deps.TokenManager))
		r.Get("/", get.New(useCase))
		r.Patch("/", update.New(useCase))
		r.Get("/courses", courses.New(useCase))

		r.Group(func(r chi.Router) {
			r.Use(user.IsReal())
			r.Patch("/password", updatepass.New(useCase))
		})
	})
}
