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

type remover interface {
	GetGroup(ctx context.Context, chatId uuid.UUID) (*entity.Group, error)
	RemoveMember(ctx context.Context, groupId, userId uuid.UUID) error
}

// @Tags group
// @Security ApiKeyAuth
// @Accept       json
// @Param request body GroupMember true "Данные о юзере"
// @Success      200   {object} model.Group      "removed"
// @Router       /v1/group/member [delete]
func NewRemove(remover remover) http.HandlerFunc {
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

		err = remover.RemoveMember(ctx, chatId, userId)
		if err != nil {
			resp.Error(r, w, err)
			return
		}

		group, _ := remover.GetGroup(ctx, chatId)

		c := model.NewGroup(group)
		render.Status(r, http.StatusOK)
		render.JSON(w, r, c)
	}
}
