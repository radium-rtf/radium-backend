package update

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/radium/entity"
	"github.com/radium-rtf/radium-backend/internal/radium/model"
	"github.com/radium-rtf/radium-backend/pkg/decode"
	"github.com/radium-rtf/radium-backend/pkg/resp"
)

type updater interface {
	UpdateGroup(ctx context.Context, group *entity.Group) (*entity.Group, error)
}

// @Tags group
// @Security ApiKeyAuth
// @Param       request body Group true " "
// @Success      201   {string} model.Group "patched"
// @Router       /v1/group/{id} [patch]
func New(updater updater) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			request Group
			ctx     = r.Context()
		)

		err := decode.Json(r.Body, &request)

		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, err.Error())
			return
		}

		id, err := uuid.Parse(chi.URLParam(r, "id"))

		if err != nil {
			resp.Error(r, w, err)
			return
		}

		group := request.toGroup(id)
		result, err := updater.UpdateGroup(ctx, group)

		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, err.Error())
			return
		}

		render.Status(r, http.StatusOK)
		render.JSON(w, r, model.NewGroup(result))
	}
}
