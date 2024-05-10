package search

import (
	"context"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/radium/entity"
	"github.com/radium-rtf/radium-backend/internal/radium/model"
	"github.com/radium-rtf/radium-backend/pkg/resp"
	"net/http"
)

type search interface {
	Search(ctx context.Context, name string, limit int) ([]*entity.Course, error)
}

// @Tags course
// @Security ApiKeyAuth
// @Param        query   query      string  true  "query"
// @Success      200   {object}  model.Card        " "
// @Router       /v1/course/search [get]
func New(search search) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			ctx    = r.Context()
			userId = ctx.Value("userId").(uuid.UUID)
		)

		query := r.URL.Query().Get("query")

		course, err := search.Search(ctx, query, 16)
		if err != nil {
			resp.Error(r, w, err)
			return
		}

		cards := model.NewCards(course, model.None, userId)
		render.Status(r, http.StatusOK)
		render.JSON(w, r, cards)
	}
}
