package update

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

type updater interface {
	Update(ctx context.Context, section *entity.Section, userId uuid.UUID) (*entity.Section, error)
}

type answerGetter interface {
	GetByUserIdAndSectionId(ctx context.Context, userId, sectionId uuid.UUID) (*entity.AnswersCollection, error)
}

// @Tags section
// @Security ApiKeyAuth
// @Param        id   path      string  true  "id"
// @Param request body Section true "Информация для раздела"
// @Success      200   {object} model.Section       "updated"
// @Router       /v1/section/{id} [put]
func New(updater updater, answerGetter answerGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			request Section
			ctx     = r.Context()
			userId  = ctx.Value("userId").(uuid.UUID)
		)

		sectionId, err := uuid.Parse(chi.URLParam(r, "id"))
		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, errors.Wrap(err, "parse id").Error())
			return
		}

		err = decode.Json(r.Body, &request)
		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, err.Error())
			return
		}

		section, err := request.toSection(sectionId)
		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, err.Error())
			return
		}

		section, err = updater.Update(ctx, section, userId)
		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, err.Error())
			return
		}

		answers, err := answerGetter.GetByUserIdAndSectionId(ctx, userId, sectionId)
		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, err.Error())
			return
		}

		dto, _, _ := model.NewSections([]*entity.Section{section}, answers.AnswerBySectionId)
		render.Status(r, http.StatusOK)
		render.JSON(w, r, dto[0])
	}
}
