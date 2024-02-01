package get

import (
	"context"
	"github.com/go-chi/render"
	"github.com/radium-rtf/radium-backend/internal/radium/entity"
	"github.com/radium-rtf/radium-backend/internal/radium/model"
	"github.com/radium-rtf/radium-backend/pkg/resp"
	"net/http"
)

type getter interface {
	Get(ctx context.Context) ([]*entity.Group, error)
}

// @Tags group
// @Security ApiKeyAuth
// @Success      200   {string}  model.Group         " "
// @Router       /v1/group [get]
func New(getter getter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var ctx = r.Context()

		groups, err := getter.Get(ctx)
		if err != nil {
			resp.Error(r, w, err)
			return
		}

		dto := model.NewGroups(groups)
		render.Status(r, http.StatusOK)
		render.JSON(w, r, dto)
	}
}
