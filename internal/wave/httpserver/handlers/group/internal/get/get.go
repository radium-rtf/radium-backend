package get

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/wave/entity"
	"github.com/radium-rtf/radium-backend/internal/wave/model"
)

type (
	getter interface {
		GetGroup(ctx context.Context, chatId uuid.UUID) (*entity.Group, error)
	}
)

// @Tags group
// @Security ApiKeyAuth
// @Param        chatId   path      string  true  "ID группы"
// @Success      200   {object} model.Group        " "
// @Router       /v1/group/{chatId} [get]
func New(getter getter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			ctx = r.Context()
		)

		chatId, err := uuid.Parse(chi.URLParam(r, "chatId"))
		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, err)
			return
		}

		group, err := getter.GetGroup(ctx, chatId)

		if err != nil {
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, err)
			return
		}

		render.Status(r, http.StatusOK)
		render.JSON(w, r, model.NewGroup(group))
	}
}
