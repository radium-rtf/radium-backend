package v1

import (
	"github.com/go-chi/chi/v5"
	"github.com/radium-rtf/radium-backend/config"
	"github.com/radium-rtf/radium-backend/internal/usecase"
	"github.com/radium-rtf/radium-backend/pkg/postgres"
)

// NewRouter
// @title       без юлерна
// @version     1.0
// @host        localhost:8080
// @BasePath    /v1
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func NewRouter(h *chi.Mux, pg *postgres.Postgres, cfg *config.Config) {
	authUseCase := usecase.NewAuthUseCase(pg, cfg)
	accountUseCase := usecase.NewAccountUseCase(pg, cfg)
	h.Route("/v1", func(r chi.Router) {
		newAuthRoutes(r, authUseCase)
		newAccountRoutes(r, accountUseCase, cfg.SigningKey)
	})
}
