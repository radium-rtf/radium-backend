package getbyslug

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/internal/lib/resp"
	"github.com/radium-rtf/radium-backend/internal/model"
	"net/http"
)

type pageGetter interface {
	GetBySlug(ctx context.Context, slug string) (*entity.Page, error)
	GetNextAndPrevious(ctx context.Context, page *entity.Page) (*model.NextAndPreviousPage, error)
	GetByIdWithUserAnswers(ctx context.Context, id, userId uuid.UUID) (*entity.Page, error)
}

// @Tags page
// @Security ApiKeyAuth
// @Param        slug   path     string  true  "slug"
// @Success 200 {object} model.Page "ok"
// @Router /v1/page/slug/{slug} [get]
func New(pageGetter pageGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var ctx = r.Context()
		userId, ok := ctx.Value("userId").(uuid.UUID)

		slug := chi.URLParam(r, "slug")

		page, err := pageGetter.GetBySlug(ctx, slug)
		if err != nil {
			resp.Error(r, w, err)
			return
		}

		nextAnsPrevious, err := pageGetter.GetNextAndPrevious(ctx, page)
		if err != nil {
			resp.Error(r, w, err)
			return
		}

		if !ok {
			render.JSON(w, r, model.NewPage(page, map[uuid.UUID][]*entity.Answer{}, nextAnsPrevious))
			return
		}

		page, err = pageGetter.GetByIdWithUserAnswers(ctx, page.Id, userId)
		if err != nil {
			resp.Error(r, w, err)
			return
		}

		dto := model.NewPage(page, map[uuid.UUID][]*entity.Answer{}, nextAnsPrevious)
		render.Status(r, http.StatusOK)
		render.JSON(w, r, dto)

	}
}
