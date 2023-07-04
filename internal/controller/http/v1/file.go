package v1

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/internal/usecase"
)

type fileRoutes struct {
	uc usecase.FileUseCase
}

func newFileRoutes(useCase usecase.FileUseCase, signingKey string) *chi.Mux {
	routes := fileRoutes{uc: useCase}

	r := chi.NewRouter()
	r.Use(authRequired(signingKey))
	r.Post("/", handler(routes.uploadFile).HTTP)
	return r
}

// @Tags file
// @Security ApiKeyAuth
// @Accept multipart/form-data
// @Param file formData file true "file"
// @Success 201 {object} entity.FileDto "created"
// @Router /upload [post]
func (r fileRoutes) uploadFile(w http.ResponseWriter, request *http.Request) *appError {
	err := request.ParseMultipartForm(10 << 20)
	if err != nil {
		return newAppError(err, http.StatusBadRequest)
	}

	file, header, err := request.FormFile("file")
	if err != nil {
		return newAppError(err, http.StatusBadRequest)
	}

	defer file.Close()
	fileRequest := entity.FileUploadRequest{File: file, Header: header}

	upload, err := r.uc.UploadFile(request.Context(), fileRequest)
	if err != nil {
		return newAppError(err, http.StatusBadRequest)
	}

	render.Status(request, http.StatusCreated)
	render.JSON(w, request, upload)

	return nil
}
