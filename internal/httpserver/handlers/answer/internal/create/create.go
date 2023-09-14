package create

import (
	"context"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/lib/decode"
	"net/http"

	"github.com/radium-rtf/radium-backend/internal/entity"
)

type creator interface {
	Create(ctx context.Context, answer *entity.Answer) (*entity.Answer, error)
}

// @Tags answer
// @Security ApiKeyAuth
// @Param       request body Answer true " "
// @Success      201   {object} verdict.Verdict "ok"
// @Router      /v1/answer [post]
func New(creator creator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			request Answer
			ctx     = r.Context()
			userId  = ctx.Value("userId").(uuid.UUID)
		)

		if err := decode.Json(r.Body, &request); err != nil {
			render.Status(r, http.StatusCreated)
			render.JSON(w, r, err.Error())
			return
		}

		answer := request.ToAnswer(userId)
		verdict, err := creator.Create(r.Context(), answer)
		if err != nil {
			render.Status(r, http.StatusCreated)
			render.JSON(w, r, err.Error())
			return
		}

		render.Status(r, http.StatusCreated)
		render.JSON(w, r, verdict)
	}
}
