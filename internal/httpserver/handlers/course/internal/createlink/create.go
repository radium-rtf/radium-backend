package createlink

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/internal/lib/decode"
	"github.com/radium-rtf/radium-backend/internal/model"
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
// @Success      201   {object} model.Link      "created"
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
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, err.Error())
			return
		}

		courseId, err := uuid.Parse(chi.URLParam(r, "courseId"))
		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, errors.Wrap(err, "parse id").Error())
			return
		}

		link := request.toLink(courseId)
		link, err = creator.CreateLink(ctx, link, userId)
		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, err.Error())
			return
		}

		c := model.NewLink(link)
		render.Status(r, http.StatusCreated)
		render.JSON(w, r, c)
	}
}
