package create

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

type creator interface {
	CreateDialogue(ctx context.Context, userId uuid.UUID, recipientId uuid.UUID) (*entity.Dialogue, error)
}

// @Tags dialogue
// @Security ApiKeyAuth
// @Accept       json
// @Param request body DialogueCreate true "Данные о реципиенте"
// @Success      201   {object} model.Dialogue      "created"
// @Success      409   {object} model.Dialogue      "already exists"
// @Router       /v1/dialogue [post]
func New(creator creator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			request    DialogueCreate
			ctx        = r.Context()
			userId, ok = ctx.Value("userId").(uuid.UUID)
		)
		if !ok {
			// resp.Error(r, w, resp.ErrUnauthorized)
			// return
			userId = uuid.Nil
		}

		err := decode.Json(r.Body, &request)
		if err != nil {
			resp.Error(r, w, err)
			return
		}

		recipientId := request.GetId()

		dialogue, err := creator.CreateDialogue(ctx, userId, recipientId)
		if err != nil {
			render.Status(r, http.StatusConflict)
			resp.Error(r, w, err)
			return
		}

		c := model.NewDialogue(dialogue)
		render.Status(r, http.StatusCreated) // or exists
		render.JSON(w, r, c)
	}
}
