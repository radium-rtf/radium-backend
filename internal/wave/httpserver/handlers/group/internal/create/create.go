package create

import (
	"context"
	"net/http"

	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/wave/entity"
	"github.com/radium-rtf/radium-backend/internal/wave/model"
	"github.com/radium-rtf/radium-backend/pkg/decode"
	"github.com/radium-rtf/radium-backend/pkg/resp"
)

type creatorGroup interface {
	CreateGroup(ctx context.Context, userId uuid.UUID, name string) (*entity.Group, error)
	AddMember(ctx context.Context, groupId, userId uuid.UUID, admin bool) error
}

// @Tags group
// @Security ApiKeyAuth
// @Accept       json
// @Param request body GroupCreate true "Данные о группе"
// @Success      201   {object} model.Dialogue      "created"
// @Router       /v1/group/create [post]
func NewGroup(creator creatorGroup) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			request GroupCreate
			ctx     = r.Context()
			userId  = ctx.Value("userId").(uuid.UUID)
		)

		err := decode.Json(r.Body, &request)
		if err != nil {
			resp.Error(r, w, err)
			return
		}

		group, err := creator.CreateGroup(ctx, userId, request.Name)
		if err != nil {
			resp.Error(r, w, err)
			return
		}

		err = creator.AddMember(ctx, group.Id, userId, true)
		if err != nil {
			resp.Error(r, w, err)
			return
		}

		c := model.NewGroup(group)
		render.Status(r, http.StatusCreated)
		render.JSON(w, r, c)
	}
}
