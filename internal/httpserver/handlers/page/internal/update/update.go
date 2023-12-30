package update

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/internal/lib/decode"
	"github.com/radium-rtf/radium-backend/internal/lib/resp"
	"github.com/radium-rtf/radium-backend/internal/model"
	"net/http"
)

type updater interface {
	Update(ctx context.Context, page *entity.Page, userId uuid.UUID) (*entity.Page, error)
}

// @Tags page
// @Security ApiKeyAuth
// @Param        pageId   path      string  true  "id"
// @Param       request body Module true " "
// @Success      201   {object} model.Page       "created"
// @Router       /v1/page/{pageId} [put]
func New(updater updater) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			request Page
			ctx     = r.Context()
			userId  = ctx.Value("userId").(uuid.UUID)
		)

		pageId, err := uuid.Parse(chi.URLParam(r, "pageId"))
		if err != nil {
			resp.Error(r, w, err)
			return
		}

		err = decode.Json(r.Body, &request)
		if err != nil {
			resp.Error(r, w, err)
			return
		}

		page := request.toPage(pageId)
		page, err = updater.Update(ctx, page, userId)
		if err != nil {
			resp.Error(r, w, err)
			return
		}

		dto := model.NewPage(page, map[uuid.UUID][]*entity.Answer{}, nil)
		render.Status(r, http.StatusCreated)
		render.JSON(w, r, dto)
	}
}
