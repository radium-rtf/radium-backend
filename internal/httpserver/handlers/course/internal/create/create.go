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

type creator interface {
	Create(ctx context.Context, course *entity.Course, editorId uuid.UUID) (*entity.Course, error)
}

// @Tags course
// @Security ApiKeyAuth
// @Accept       json
// @Param request body Course true "Данные о курсе"
// @Success      201   {object} model.Course       "created"
// @Router       /v1/course [post]
func New(creator creator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			request Course
			ctx     = r.Context()
			userId  = ctx.Value("userId").(uuid.UUID)
		)

		err := decode.Json(r.Body, &request)
		if err != nil {
			resp.Error(r, w, err)
			return
		}

		course := request.toCourse(userId)
		course, err = creator.Create(ctx, course, userId)
		if err != nil {
			resp.Error(r, w, err)
			return
		}

		c := model.NewCourse(course, map[uuid.UUID][]*entity.Answer{}, userId)
		render.Status(r, http.StatusCreated)
		render.JSON(w, r, c)
	}
}
