package handlers

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/httprate"
	_ "github.com/radium-rtf/radium-backend/docs/wave"
	"github.com/radium-rtf/radium-backend/internal/wave/usecase"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
	"time"
)

// @title       wave
// @version     1.0
// @BasePath    /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func NewRouter(h *chi.Mux, useCases usecase.UseCases) {
	log := newLogger()

	h.Use(httprate.Limit(100, time.Second))
	h.Use(cors.AllowAll().Handler)
	h.Use(middleware.Recoverer)
	h.Use(middleware.RequestID)
	h.Use(newHandlerLogger(log))

	h.Get("/healthz", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
	})

	swaggerHandler := httpSwagger.Handler(
		httpSwagger.URL("/wave/doc.json"),
		httpSwagger.InstanceName("wave"))
	h.Get("/wave/*", swaggerHandler)

	for _, pattern := range []string{"/", "/swagger/*"} {
		h.Get(pattern, func(writer http.ResponseWriter, request *http.Request) {
			writer.Header().Add("Location", "/wave/index.html")
			writer.WriteHeader(http.StatusTemporaryRedirect)
		})
	}
}
