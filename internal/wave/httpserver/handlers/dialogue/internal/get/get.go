package get

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/wave/entity"
	"github.com/radium-rtf/radium-backend/internal/wave/model"
	"github.com/radium-rtf/radium-backend/pkg/resp"
)

type getter interface {
	GetDialogue(ctx context.Context, chatId uuid.UUID) (*entity.Dialogue, error)
}

// @Tags message
// @Security ApiKeyAuth
// @Param        chatId   path      string  true  "ID группы/диалога"
// @Success      200   {object} []model.Dialogue        " "
// @Router       /v1/dialogue/{chatId} [get]
func New(getter getter) http.HandlerFunc {
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

		dialogue, err := getter.GetDialogue(ctx, chatId)
		if err != nil {
			resp.Error(r, w, err)
			return
		}

		c := model.NewDialogue(dialogue)

		render.Status(r, http.StatusOK)
		render.JSON(w, r, c)
	}
}
