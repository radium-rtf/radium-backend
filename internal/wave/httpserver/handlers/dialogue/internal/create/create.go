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

type creatorDialogue interface {
	CreateDialogue(ctx context.Context, dialogue *entity.Dialogue) error
	GetDialogueByUsers(ctx context.Context, firstUser, secondUser uuid.UUID) (*entity.Dialogue, error)
}

// @Tags dialogue
// @Security ApiKeyAuth
// @Accept       json
// @Param request body DialogueCreate true "Данные о реципиенте"
// @Success      201   {object} model.Dialogue      "created"
// @Failure      409   {object} model.Dialogue      "already exists"
// @Router       /v1/dialogue/create [post]
func NewDialogue(creator creatorDialogue) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			request DialogueCreate
			ctx     = r.Context()
			userId  = ctx.Value("userId").(uuid.UUID)
		)

		err := decode.Json(r.Body, &request)
		if err != nil {
			resp.Error(r, w, err)
			return
		}

		recipientId := request.UserId

		existingDialogue, err := creator.GetDialogueByUsers(ctx, userId, recipientId)
		if err == nil {
			c := model.NewDialogue(existingDialogue)
			render.Status(r, http.StatusConflict)
			render.JSON(w, r, c)
			return
		}

		dialogue := &entity.Dialogue{
			Id:           uuid.New(),
			FirstUserId:  userId,
			SecondUserId: recipientId,
		}
		err = creator.CreateDialogue(ctx, dialogue)
		if err != nil {
			resp.Error(r, w, err)
			return
		}

		c := model.NewDialogue(dialogue)
		render.Status(r, http.StatusCreated) // or exists
		render.JSON(w, r, c)
	}
}
