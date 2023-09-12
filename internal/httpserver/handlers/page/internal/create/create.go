package create

import (
	"context"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/internal/model"
	"github.com/radium-rtf/radium-backend/pkg/validator"
	"net/http"
)

type creator interface {
	Create(ctx context.Context, page *entity.Page) (*entity.Page, error)
}

// @Tags page
// @Security ApiKeyAuth
// @Param       request body Page true "создание"
// @Success      201   {object} model.Page "ok"
// @Router      /v1/page [post]
func New(creator creator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			request Page
			ctx     = r.Context()
		)

		err := render.DecodeJSON(r.Body, &request)
		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, err.Error())
			return
		}

		err = validator.Struct(request)
		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, err.Error())
			return
		}

		page := request.toPage()
		page, err = creator.Create(ctx, page)
		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, err.Error())
			return
		}

		m := model.NewPage(page, map[uuid.UUID]*entity.Answer{})
		render.Status(r, http.StatusCreated)
		render.JSON(w, r, m)
	}
}
