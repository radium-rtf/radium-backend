package create

import (
	"context"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/internal/lib/resp"
	"github.com/radium-rtf/radium-backend/internal/model"
	"github.com/radium-rtf/radium-backend/pkg/decode"
	"net/http"
)

type useCase interface {
	GetNextAndPrevious(ctx context.Context, page *entity.Page) (*model.NextAndPreviousPage, error)
	Create(ctx context.Context, page *entity.Page, editorId uuid.UUID) (*entity.Page, error)
}

// @Tags page
// @Security ApiKeyAuth
// @Param       request body Page true "создание"
// @Success      201   {object} model.Page "ok"
// @Router      /v1/page [post]
func New(useCase useCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			request Page
			ctx     = r.Context()
			userId  = r.Context().Value("userId").(uuid.UUID)
		)

		err := decode.Json(r.Body, &request)
		if err != nil {
			resp.Error(r, w, err)
			return
		}

		page := request.toPage()
		page, err = useCase.Create(ctx, page, userId)
		if err != nil {
			resp.Error(r, w, err)
			return
		}

		nextAnsPrevious, err := useCase.GetNextAndPrevious(ctx, page)
		if err != nil {
			resp.Error(r, w, err)
			return
		}

		m := model.NewPage(page, map[uuid.UUID][]*entity.Answer{}, nextAnsPrevious)
		render.Status(r, http.StatusCreated)
		render.JSON(w, r, m)
	}
}
