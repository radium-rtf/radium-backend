package getbyslug

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	entity2 "github.com/radium-rtf/radium-backend/internal/radium/entity"
	"github.com/radium-rtf/radium-backend/internal/radium/model"
	"github.com/radium-rtf/radium-backend/pkg/resp"
	"net/http"
)

type pageGetter interface {
	GetBySlug(ctx context.Context, slug string) (*entity2.Page, error)
	GetNextAndPrevious(ctx context.Context, page *entity2.Page) (*model.NextAndPreviousPage, error)
	GetByIdWithUserAnswers(ctx context.Context, id, userId uuid.UUID) (*entity2.Page, error)
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
			render.JSON(w, r, model.NewPage(page, map[uuid.UUID][]*entity2.Answer{}, nextAnsPrevious))
			return
		}

		page, err = pageGetter.GetByIdWithUserAnswers(ctx, page.Id, userId)
		if err != nil {
			resp.Error(r, w, err)
			return
		}

		dto := model.NewPage(page, map[uuid.UUID][]*entity2.Answer{}, nextAnsPrevious)
		render.Status(r, http.StatusOK)
		render.JSON(w, r, dto)

	}
}
