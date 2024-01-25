package create

import (
	"context"
	"github.com/go-chi/render"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/internal/lib/resp"
	"github.com/radium-rtf/radium-backend/internal/model"
	"github.com/radium-rtf/radium-backend/pkg/decode"
	"net/http"
)

type creator interface {
	Create(ctx context.Context, group *entity.Group) (*entity.Group, error)
}

// @Tags group
// @Security ApiKeyAuth
// @Param       request body Group true " "
// @Success      201   {string} model.Group "created"
// @Router       /v1/group [post]
func New(creator creator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			request Group
			ctx     = r.Context()
		)

		err := decode.Json(r.Body, &request)
		if err != nil {
			resp.Error(r, w, err)
			return
		}

		group := request.toGroup()
		group, err = creator.Create(ctx, group)
		if err != nil {
			resp.Error(r, w, err)
			return
		}

		dto := model.NewGroup(group)
		render.Status(r, http.StatusCreated)
		render.JSON(w, r, dto)
	}
}
