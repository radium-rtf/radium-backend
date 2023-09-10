package app

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/radium-rtf/radium-backend/config"
	"github.com/radium-rtf/radium-backend/internal/httpserver/handlers"
	"github.com/radium-rtf/radium-backend/internal/usecase"
	"github.com/radium-rtf/radium-backend/pkg/httpserver"
	httpSwagger "github.com/swaggo/http-swagger"
)

func startHttpServer(http config.HTTP, dependencies usecase.Dependencies) *httpserver.Server {
	r := chi.NewRouter()
	r.Use(cors.AllowAll().Handler)
	r.Get("/swagger/*", httpSwagger.Handler(httpSwagger.URL("/swagger/doc.json")))

	useCases := usecase.NewUseCases(dependencies)
	handlers.NewRouter(r, useCases)

	return httpserver.New(r,
		httpserver.Port(http.Port),
		httpserver.MaxHeaderBytes(http.MaxHeaderBytes),
		httpserver.IdleTimeout(http.IdleTimeout),
		httpserver.WriteTimeout(http.WriteTimeout),
		httpserver.ReadTimeout(http.ReadTimeout),
	)
}
