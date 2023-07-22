package auth

import (
	"github.com/go-chi/chi/v5"
	"github.com/radium-rtf/radium-backend/internal/httpserver/handlers/auth/internal/refresh"
	"github.com/radium-rtf/radium-backend/internal/httpserver/handlers/auth/internal/signin"
	"github.com/radium-rtf/radium-backend/internal/httpserver/handlers/auth/internal/signup"
	"github.com/radium-rtf/radium-backend/internal/lib/session"
	"github.com/radium-rtf/radium-backend/internal/usecase"
	"github.com/radium-rtf/radium-backend/internal/usecase/repo/postgres"
	"github.com/radium-rtf/radium-backend/pkg/hash"
	"github.com/radium-rtf/radium-backend/pkg/postgres/db"
)

func New(r *chi.Mux, pg *db.Query, hasher hash.Hasher, session session.Session) {
	sessionRepo := postgres.NewSessionRepo(pg)
	userRepo := postgres.NewUserRepo(pg)
	useCase := usecase.NewAuthUseCase(userRepo, sessionRepo, hasher, session)

	r.Route("/v1/auth", func(r chi.Router) {
		r.Post("/signin", signin.New(useCase))
		r.Post("/signup", signup.New(useCase))
		r.Post("/refresh", refresh.New(useCase))
	})
}
