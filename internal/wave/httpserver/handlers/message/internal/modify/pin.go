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

type pinner interface {
	PinMessage(ctx context.Context, userId, messageId uuid.UUID, status bool) (*model.Message, error)
}

// @Tags message
// @Security ApiKeyAuth
// @Accept       json
// @Param request body MessageGeneric true "Сообщение для закрепления"
// @Success      200   {object} model.Message        "(un)pinned"
// @Failure      404
// @Router       /v1/message/pin [PATCH]
// @Router       /v1/message/pin [DELETE]
func NewPin(pinner pinner) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			request MessageGeneric
			ctx     = r.Context()
			userId  = ctx.Value("userId").(uuid.UUID)
		)

		err := decode.Json(r.Body, &request)
		if err != nil {
			resp.Error(r, w, err)
			return
		}

		messageId := request.MessageId

		var message *model.Message
		if r.Method == http.MethodPatch {
			message, err = pinner.PinMessage(ctx, userId, messageId, true)
		} else {
			message, err = pinner.PinMessage(ctx, userId, messageId, false)
		}

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
