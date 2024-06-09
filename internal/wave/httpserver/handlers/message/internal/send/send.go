package send

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

type sender interface {
	SendMessage(ctx context.Context, message *entity.Message, chatId uuid.UUID) (*model.Message, error)
}

// @Tags message
// @Security ApiKeyAuth
// @Accept       json
// @Param request body MessageSend true "Сообщение и направление"
// @Success      201   {object} model.Message        "sent"
// @Router       /v1/message [post]
func New(sender sender) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			request MessageSend
			ctx     = r.Context()
			userId  = ctx.Value("userId").(uuid.UUID)
		)

		err := decode.Json(r.Body, &request)
		if err != nil {
			resp.Error(r, w, err)
			return
		}

		chatId, content := request.ChatId, request.Content

		contentObject := &entity.Content{
			DBModel: entity.DBModel{
				Id: uuid.New(),
			},
			Text: content.Text,
		}
		messageObject := &entity.Message{
			DBModel: entity.DBModel{
				Id: uuid.New(),
			},
			SenderId:  userId,
			ContentId: contentObject.Id,
			Content:   contentObject,
		}

		message, err := sender.SendMessage(ctx, messageObject, chatId)
		if err != nil {
			resp.Error(r, w, err)
			return
		}

		c := message
		render.Status(r, http.StatusCreated)
		render.JSON(w, r, c)
	}
}
