package getbyid

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/internal/lib/resp"
	"github.com/radium-rtf/radium-backend/internal/model"
	"net/http"
)

type getter interface {
	GetById(ctx context.Context, id uuid.UUID) (*entity.Group, error)
}

// @Tags group
// @Security ApiKeyAuth
// @Param        id   path      string  true  "groupId"
// @Success      200   {string}  model.Group         " "
// @Router       /v1/group/{id} [get]
func New(getter getter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var ctx = r.Context()

		id, err := uuid.Parse(chi.URLParam(r, "id"))
		if err != nil {
			resp.Error(r, w, err)
			return
		}

		group, err := getter.GetById(ctx, id)
		if err != nil {
			resp.Error(r, w, err)
			return
		}

		dto := model.NewGroup(group)
		render.Status(r, http.StatusOK)
		render.JSON(w, r, dto)
	}
}
