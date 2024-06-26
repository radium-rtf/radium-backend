package createlink

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/radium/entity"
	"github.com/radium-rtf/radium-backend/internal/radium/model"
	"github.com/radium-rtf/radium-backend/pkg/decode"
	"github.com/radium-rtf/radium-backend/pkg/resp"
	"net/http"
)

type creator interface {
	CreateLink(ctx context.Context, link *entity.Link, editorId uuid.UUID) (*entity.Link, error)
}

// @Tags course-link
// @Security ApiKeyAuth
// @Accept       json
// @Param request body Link true "Данные о ссылке"
// @Param        courseId   path      string  true  "id"
// @Success      201   {object} createlink.Link      "created"
// @Router       /v1/course/link/{courseId} [post]
func New(creator creator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			request Link
			ctx     = r.Context()
			userId  = ctx.Value("userId").(uuid.UUID)
		)

		err := decode.Json(r.Body, &request)
		if err != nil {
			resp.Error(r, w, err)
			return
		}

		courseId, err := uuid.Parse(chi.URLParam(r, "courseId"))
		if err != nil {
			resp.Error(r, w, err)
			return
		}

		link := request.ToLink(courseId)
		link, err = creator.CreateLink(ctx, link, userId)
		if err != nil {
			resp.Error(r, w, err)
			return
		}

		c := model.NewLink(link)
		render.Status(r, http.StatusCreated)
		render.JSON(w, r, c)
	}
}
