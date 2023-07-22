package create

import (
	"context"
	"encoding/json"
	"github.com/go-chi/render"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/internal/model"
	"net/http"
)

type creator interface {
	Create(ctx context.Context, course *entity.Course) (*entity.Course, error)
}

// @Tags course
// @Security ApiKeyAuth
// @Accept       json
// @Param request body Request true "Данные о курсе"
// @Success      201   {object} model.Course       "created"
// @Router       /v1/course [post]
func New(creator creator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			request Request
			ctx     = r.Context()
		)

		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			render.Status(r, http.StatusCreated)
			render.JSON(w, r, err.Error())
			return
		}

		course := request.ToCourse()
		course, err = creator.Create(ctx, course)
		if err != nil {
			render.Status(r, http.StatusCreated)
			render.JSON(w, r, err.Error())
			return
		}

		c := model.NewCourse(course)
		render.Status(r, http.StatusCreated)
		render.JSON(w, r, c)
	}
}
