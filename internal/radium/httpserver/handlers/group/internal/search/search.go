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
	Search(ctx context.Context, name string, limit int) ([]*entity.Group, error)
}

// @Tags group
// @Security ApiKeyAuth
// @Param        query   query      string  true  "query"
// @Success      200   {object}  model.Group        " "
// @Router       /v1/group/search [get]
func New(search search) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			ctx = r.Context()
			_   = ctx.Value("userId").(uuid.UUID)
		)

		query := r.URL.Query().Get("query")

		groups, err := search.Search(ctx, query, 16)
		if err != nil {
			resp.Error(r, w, err)
			return
		}

		render.Status(r, http.StatusOK)
		render.JSON(w, r, model.NewGroups(groups))
	}
}
