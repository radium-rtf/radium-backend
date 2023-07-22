package getbyid

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/model"
	"net/http"
)

type getter interface {
	GetById(ctx context.Context, id uuid.UUID, userId *uuid.UUID) (*model.Page, error)
}

// @Tags page
// @Security ApiKeyAuth
// @Param        id   path     string  true  "page id"
// @Success 200 {object} model.Page "ok"
// @Router /v1/page/{id} [get]
func New(getter getter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var ctx = r.Context()

		id, err := uuid.Parse(chi.URLParam(r, "id"))
		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, err.Error())
			return
		}

		userId, ok := ctx.Value("userId").(uuid.UUID)
		var page *model.Page
		if !ok {
			page, err = getter.GetById(ctx, id, nil)
		} else {
			page, err = getter.GetById(ctx, id, &userId)
		}

		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, err.Error())
			return
		}

		render.Status(r, http.StatusOK)
		render.JSON(w, r, page)
	}
}
