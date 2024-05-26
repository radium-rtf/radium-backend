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

type editor interface {
	EditMessage(ctx context.Context, userId, messageId uuid.UUID, content model.Content) (*model.Message, error)
}

// @Tags message
// @Security ApiKeyAuth
// @Accept       json
// @Param request body MessageEdit true "Сообщение"
// @Success      200   {object} model.Message        "edited"
// @Failure      404
// @Router       /v1/message [patch]
func NewEdit(editor editor) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			request MessageEdit
			ctx     = r.Context()
			userId  = ctx.Value("userId").(uuid.UUID)
		)

		err := decode.Json(r.Body, &request)
		if err != nil {
			resp.Error(r, w, err)
			return
		}

		message, err := editor.EditMessage(ctx, userId, request.MessageId, request.Content)
		if err != nil {
			resp.Error(r, w, err)
			return
		}

		c := message
		render.Status(r, http.StatusOK)
		render.JSON(w, r, c)
	}
}
