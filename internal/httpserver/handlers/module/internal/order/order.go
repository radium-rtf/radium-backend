package order

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/internal/lib/decode"
	"github.com/radium-rtf/radium-backend/internal/model"
	"net/http"
)

type updater interface {
	UpdateOrder(ctx context.Context, id, editorId uuid.UUID, order uint) (*entity.Module, error)
}

// @Tags module
// @Security ApiKeyAuth
// @Accept json
// @Param        id   path     string  true  "module id"
// @Param request body Order true " "
// @Success 200 {object} model.Module " "
// @Router /v1/module/{id}/order [patch]
func New(updater updater) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			order  Order
			ctx    = r.Context()
			userId = r.Context().Value("userId").(uuid.UUID)
		)

		err := decode.Json(r.Body, &order)
		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, err.Error())
			return
		}

		id, err := uuid.Parse(chi.URLParam(r, "id"))
		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, errors.Wrap(err, "parse id").Error())
			return
		}

		module, err := updater.UpdateOrder(ctx, id, userId, order.Order)
		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, err.Error())
			return
		}

		dto := model.NewModule(module, map[uuid.UUID][]*entity.Answer{})
		render.Status(r, http.StatusOK)
		render.JSON(w, r, dto)
	}
}
