package order

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/internal/lib/resp"
	"github.com/radium-rtf/radium-backend/internal/model"
	"github.com/radium-rtf/radium-backend/pkg/decode"
	"net/http"
)

type updater interface {
	UpdateOrder(ctx context.Context, id, editorId uuid.UUID, order uint) (*entity.Section, error)
}

// @Tags section
// @Security ApiKeyAuth
// @Accept json
// @Param        id   path     string  true  "section id"
// @Param request body Order true " "
// @Success 200 {object} model.Section " "
// @Router /v1/section/{id}/order [patch]
func New(updater updater) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			order  Order
			ctx    = r.Context()
			userId = r.Context().Value("userId").(uuid.UUID)
		)

		err := decode.Json(r.Body, &order)
		if err != nil {
			resp.Error(r, w, err)
			return
		}

		id, err := uuid.Parse(chi.URLParam(r, "id"))
		if err != nil {
			resp.Error(r, w, err)
			return
		}

		section, err := updater.UpdateOrder(ctx, id, userId, order.Order)
		if err != nil {
			resp.Error(r, w, err)
			return
		}

		dto := model.NewSection(section, nil, 0)
		render.Status(r, http.StatusOK)
		render.JSON(w, r, dto)
	}
}
