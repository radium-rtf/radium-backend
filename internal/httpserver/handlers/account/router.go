package account

import (
	"github.com/go-chi/chi/v5"
	"github.com/radium-rtf/radium-backend/internal/httpserver/handlers/account/internal/courses"
	"github.com/radium-rtf/radium-backend/internal/httpserver/handlers/account/internal/get"
	"github.com/radium-rtf/radium-backend/internal/httpserver/handlers/account/internal/update"
	"github.com/radium-rtf/radium-backend/internal/httpserver/handlers/account/internal/updatepass"
	mwAuth "github.com/radium-rtf/radium-backend/internal/httpserver/middleware/auth"
	"github.com/radium-rtf/radium-backend/internal/usecase/repo/postgres"
	"github.com/radium-rtf/radium-backend/pkg/hash"
	"github.com/radium-rtf/radium-backend/pkg/postgres/db"

	"github.com/radium-rtf/radium-backend/internal/usecase"
	"github.com/radium-rtf/radium-backend/pkg/auth"
)

func New(r *chi.Mux, pg *db.Query, manager auth.TokenManager, hasher hash.Hasher) {
	userRepo := postgres.NewUserRepo(pg)
	courseRepo := postgres.NewCourseRepo(pg)
	useCase := usecase.NewAccountUseCase(userRepo, courseRepo, hasher)

	r.Route("/v1/account", func(r chi.Router) {
		r.Use(mwAuth.Required(manager))
		r.Get("/", get.New(useCase))
		r.Patch("/", update.New(useCase))
		r.Patch("/password", updatepass.New(useCase))
		r.Get("/courses", courses.New(useCase))
	})
}
