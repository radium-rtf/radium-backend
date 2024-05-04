package create

import (
	"context"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	entity2 "github.com/radium-rtf/radium-backend/internal/radium/entity"
	"github.com/radium-rtf/radium-backend/internal/radium/model"
	"github.com/radium-rtf/radium-backend/pkg/decode"
	"github.com/radium-rtf/radium-backend/pkg/resp"
	"net/http"
)

type creator interface {
	Create(ctx context.Context, course *entity2.Course, editorId uuid.UUID) (*entity2.Course, error)
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

		course.Groups = make([]*entity2.Group, 0)
		c := model.NewCourse(course, map[uuid.UUID][]*entity2.Answer{}, userId)
		render.Status(r, http.StatusCreated)
		render.JSON(w, r, c)
	}
}
