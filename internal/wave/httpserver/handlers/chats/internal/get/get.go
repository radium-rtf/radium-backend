package get

import (
	"context"
	"net/http"

	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/wave/entity"
	"github.com/radium-rtf/radium-backend/internal/wave/model"
)

type (
	getter interface {
		GetDialogues(ctx context.Context, userId uuid.UUID) ([]*entity.Dialogue, error)
	}
	messageGetter interface {
		GetLastMessage(ctx context.Context, chatId uuid.UUID) (*model.Message, error)
	}
)

// @Tags chats
// @Security ApiKeyAuth
// @Success      200   {object} []model.Chat        " "
// @Router       /v1/chats [get]
func New(getter getter, messageGetter messageGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			ctx    = r.Context()
			userId = r.Context().Value("userId").(uuid.UUID)
		)

		dialogues, err := getter.GetDialogues(ctx, userId)

		if err != nil {
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, err)
			return
		}

		c := make([]model.Chat, 0, len(dialogues))
		for _, d := range dialogues {
			message, _ := messageGetter.GetLastMessage(ctx, d.Id)

			c = append(c, model.NewChat(
				d.Id,
				d.Id.String(), // TODO: change name
				"dialogue",
				message,
			))
		}

		render.Status(r, http.StatusOK)
		render.JSON(w, r, c)
	}
}
