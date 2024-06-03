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
		GetChats(ctx context.Context, userId uuid.UUID) ([]*entity.Chat, error)
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

		chats, err := getter.GetChats(ctx, userId)

		if err != nil {
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, err)
			return
		}

		models := make([]model.Chat, 0, len(chats))
		for _, c := range chats {
			message, _ := messageGetter.GetLastMessage(ctx, c.Id)

			models = append(models, *model.NewChat(c, message))
		}

		render.Status(r, http.StatusOK)
		render.JSON(w, r, models)
	}
}
