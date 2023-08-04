package create

import (
	"context"
	"github.com/go-chi/render"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/internal/model"
	"net/http"
)

type creator interface {
	Create(ctx context.Context, group *entity.Group) (*entity.Group, error)
}

// @Tags group
// @Security ApiKeyAuth
// @Param       request body Request true " "
// @Success      201   {string} model.Group "created"
// @Router       /v1/group [post]
func New(creator creator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			request Request
			ctx     = r.Context()
		)

		err := render.DecodeJSON(r.Body, &request)
		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, err.Error())
			return
		}

		group := request.ToGroup()
		group, err = creator.Create(ctx, group)
		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, err.Error())
			return
		}

		dto := model.NewGroup(group)
		render.Status(r, http.StatusCreated)
		render.JSON(w, r, dto)
	}
}