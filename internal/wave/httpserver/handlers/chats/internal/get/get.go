package get

import (
	"context"
	"net/http"

	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/wave/entity"
)

type getter interface {
	GetDialogues(ctx context.Context, userId uuid.UUID) ([]*entity.Dialogue, error)
}

type Chat struct {
	Id   uuid.UUID `json:"id"`
	Name string    `json:"name"`
	Type string    `json:"type"`
}

// @Tags chats
// @Security ApiKeyAuth
// @Success      200   {object} []Chat        " "
// @Router       /v1/chats [get]
func New(getter getter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			ctx        = r.Context()
			userId, ok = r.Context().Value("userId").(uuid.UUID)
		)
		if !ok {
			userId = uuid.Nil
			// resp.Error(r, w, resp.ErrUnauthorized)
			// return
		}

		dialogues, err := getter.GetDialogues(ctx, userId)

		if err != nil {
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, err)
			return
		}

		c := make([]Chat, 0, len(dialogues))
		for _, d := range dialogues {
			c = append(c, Chat{
				Id:   d.Id,
				Name: d.FirstUserId.String() + " / " + d.SecondUserId.String(),
				Type: "dialogue",
			})
		}

		render.Status(r, http.StatusOK)
		render.JSON(w, r, c)
	}
}
