package answer

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/internal/lib/resp"
	"net/http"
)

type getter interface {
	GetFullSectionById(ctx context.Context, id, editorId uuid.UUID) (*entity.Section, error)
}

// @Tags section
// @Security ApiKeyAuth
// @Accept json
// @Param        id   path     string  true  "section id"
// @Success 200 {object} Answer " "
// @Router /v1/section/{id}/answer [get]
func New(getter getter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			ctx    = r.Context()
			userId = r.Context().Value("userId").(uuid.UUID)
		)

		id, err := uuid.Parse(chi.URLParam(r, "id"))
		if err != nil {
			resp.Error(r, w, err)
			return
		}

		section, err := getter.GetFullSectionById(ctx, id, userId)
		if err != nil {
			resp.Error(r, w, err)
			return
		}

		dto := Answer{Answer: section.Answer, Answers: section.Answers}
		render.Status(r, http.StatusOK)
		render.JSON(w, r, dto)
	}
}
