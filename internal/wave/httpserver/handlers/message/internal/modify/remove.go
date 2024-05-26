package modify

import (
	"context"
	"net/http"

	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/wave/model"
	"github.com/radium-rtf/radium-backend/pkg/decode"
	"github.com/radium-rtf/radium-backend/pkg/resp"
)

type remover interface {
	RemoveMessage(ctx context.Context, userId, messageId uuid.UUID) (*model.Message, error)
}

// @Tags message
// @Security ApiKeyAuth
// @Accept       json
// @Param request body MessageRemove true "Сообщение для удаления"
// @Success      200   {object} model.Message        "deleted"
// @Failure      404
// @Router       /v1/message [delete]
func NewRemove(remover remover) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			request MessageRemove
			ctx     = r.Context()
			userId  = ctx.Value("userId").(uuid.UUID)
		)

		err := decode.Json(r.Body, &request)
		if err != nil {
			resp.Error(r, w, err)
			return
		}

		messageId := request.MessageId

		message, err := remover.RemoveMessage(ctx, userId, messageId)
		if err != nil {
			resp.Error(r, w, err)
			render.Status(r, http.StatusNotFound)
			return
		}

		c := message
		render.Status(r, http.StatusOK)
		render.JSON(w, r, c)
	}
}
