package modify

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

type adder interface {
	GetGroup(ctx context.Context, chatId uuid.UUID) (*entity.Group, error)
	AddMember(ctx context.Context, groupId, userId uuid.UUID, admin bool) error
}

// @Tags group
// @Security ApiKeyAuth
// @Accept       json
// @Param request body GroupMember true "Данные о юзере"
// @Success      200   {object} model.Group      "added"
// @Router       /v1/group/member [post]
func NewAdd(adder adder) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			request GroupMember
			ctx     = r.Context()
		)

		err := decode.Json(r.Body, &request)
		if err != nil {
			resp.Error(r, w, err)
			return
		}

		userId := request.UserId
		chatId := request.ChatId

		err = adder.AddMember(ctx, chatId, userId, false)
		if err != nil {
			resp.Error(r, w, err)
			return
		}

		group, _ := adder.GetGroup(ctx, chatId)

		c := model.NewGroup(group)
		render.Status(r, http.StatusOK)
		render.JSON(w, r, c)
	}
}
