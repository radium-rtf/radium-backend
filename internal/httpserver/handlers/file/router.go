package file

import (
	"github.com/go-chi/chi/v5"
	"github.com/radium-rtf/radium-backend/internal/httpserver/handlers/file/internal/upload"
	mwAuth "github.com/radium-rtf/radium-backend/internal/httpserver/middleware/auth"
	"github.com/radium-rtf/radium-backend/internal/usecase"
	"github.com/radium-rtf/radium-backend/pkg/auth"
	"github.com/radium-rtf/radium-backend/pkg/filestorage"
)

func New(r *chi.Mux, storage filestorage.Storage, manager auth.TokenManager) {
	useCase := usecase.NewFileUseCase(storage)
	r.Group(func(r chi.Router) {
		r.Use(mwAuth.Required(manager))
		r.Post("/", upload.New(useCase))
	})
}
