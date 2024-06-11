package get

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/wave/model"
	"github.com/radium-rtf/radium-backend/pkg/resp"
)

type pinGetter interface {
	GetPinnedMessages(ctx context.Context, chatId uuid.UUID) ([]*model.Message, error)
}

// @Tags message
// @Security ApiKeyAuth
// @Param        chatId   path      string  true  "ID группы/диалога"
// @Success      200   {object} []model.Message        " "
// @Router       /v1/messages/{chatId}/pins [get]
func NewPins(getter pinGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			ctx = r.Context()
			// userId = r.Context().Value("userId").(uuid.UUID)
		)

		chatId, err := uuid.Parse(chi.URLParam(r, "chatId"))
		if err != nil {
			resp.Error(r, w, err)
			return
		}

		messages, err := getter.GetPinnedMessages(ctx, chatId)
		if err != nil {
			resp.Error(r, w, err)
			return
		}

		render.Status(r, http.StatusOK)
		render.JSON(w, r, messages)
	}
}
