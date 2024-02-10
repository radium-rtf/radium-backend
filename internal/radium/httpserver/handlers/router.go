package handlers

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/httprate"
	_ "github.com/radium-rtf/radium-backend/docs/radium"
	"github.com/radium-rtf/radium-backend/internal/radium/httpserver/handlers/account"
	"github.com/radium-rtf/radium-backend/internal/radium/httpserver/handlers/answer"
	"github.com/radium-rtf/radium-backend/internal/radium/httpserver/handlers/auth"
	"github.com/radium-rtf/radium-backend/internal/radium/httpserver/handlers/author"
	"github.com/radium-rtf/radium-backend/internal/radium/httpserver/handlers/course"
	"github.com/radium-rtf/radium-backend/internal/radium/httpserver/handlers/file"
	"github.com/radium-rtf/radium-backend/internal/radium/httpserver/handlers/group"
	"github.com/radium-rtf/radium-backend/internal/radium/httpserver/handlers/module"
	"github.com/radium-rtf/radium-backend/internal/radium/httpserver/handlers/page"
	"github.com/radium-rtf/radium-backend/internal/radium/httpserver/handlers/review"
	"github.com/radium-rtf/radium-backend/internal/radium/httpserver/handlers/role"
	"github.com/radium-rtf/radium-backend/internal/radium/httpserver/handlers/section"
	"github.com/radium-rtf/radium-backend/internal/radium/httpserver/handlers/teacher"
	"github.com/radium-rtf/radium-backend/internal/radium/usecase"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
	"time"
)

// @title       radium
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
		httpSwagger.URL("/radium/doc.json"),
		httpSwagger.InstanceName("radium"))

	h.Get("/radium/*", swaggerHandler)

	for _, pattern := range []string{"/", "/swagger/*"} {
		h.Get(pattern, func(writer http.ResponseWriter, request *http.Request) {
			writer.Header().Add("Location", "/radium/index.html")
			writer.WriteHeader(http.StatusTemporaryRedirect)
		})
	}

	h.Group(func(r chi.Router) {
		routes(r, useCases)
	})
}

func routes(h chi.Router, useCases usecase.UseCases) {
	course.New(h, useCases)
	module.New(h, useCases)
	page.New(h, useCases)
	section.New(h, useCases)

	group.New(h, useCases)
	teacher.New(h, useCases)
	review.New(h, useCases)

	answer.New(h, useCases)

	auth.New(h, useCases)
	account.New(h, useCases)
	role.New(h, useCases)

	file.New(h, useCases)

	author.New(h, useCases)
}
