package get

import (
	"context"
	"net/http"

	"github.com/go-chi/render"
	"github.com/radium-rtf/radium-backend/internal/wave/entity"
	"github.com/radium-rtf/radium-backend/internal/wave/model"
	"github.com/radium-rtf/radium-backend/pkg/resp"
)

type getter interface {
	GetDialogue(ctx context.Context) (*entity.Dialogue, error)
}

// @Tags dialogue
// @Security ApiKeyAuth
// @Success      200   {object} model.Dialogue        "ok"
// @Router       /v1/dialogue [get]
func New(getter getter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		// userId, _ := ctx.Value("userId").(uuid.UUID)

		dialogue, err := getter.GetDialogue(ctx)
		if err != nil {
			resp.Error(r, w, err)
			return
		}

		c := model.NewDialogue(dialogue)
		render.Status(r, http.StatusOK)
		render.JSON(w, r, c)
	}
}
