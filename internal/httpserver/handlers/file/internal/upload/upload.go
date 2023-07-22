package upload

import (
	"context"
	"github.com/go-chi/render"
	"github.com/radium-rtf/radium-backend/internal/model"
	"mime/multipart"
	"net/http"
)

type uploader interface {
	UploadFile(ctx context.Context, file multipart.File, header *multipart.FileHeader) (model.File, error)
}

// @Tags file
// @Security ApiKeyAuth
// @Accept multipart/form-data
// @Param file formData file true "file"
// @Success 201 {object} model.File "created"
// @Router /v1/upload [post]
func New(uploader uploader) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var ctx = r.Context()

		err := r.ParseMultipartForm(10 << 20)
		if err != nil {
			render.Status(r, http.StatusCreated)
			render.JSON(w, r, err.Error())
			return
		}

		file, header, err := r.FormFile("file")
		if err != nil {
			render.Status(r, http.StatusCreated)
			render.JSON(w, r, err.Error())
			return
		}

		defer file.Close()

		upload, err := uploader.UploadFile(ctx, file, header)
		if err != nil {
			render.Status(r, http.StatusCreated)
			render.JSON(w, r, err.Error())
			return
		}

		render.Status(r, http.StatusCreated)
		render.JSON(w, r, upload)
	}
}
