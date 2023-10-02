package getbyid

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/internal/model"
	"net/http"
)

type pageGetter interface {
	GetById(ctx context.Context, id uuid.UUID) (*entity.Page, error)
}

type answersGetter interface {
	GetBySections(ctx context.Context, ids []uuid.UUID, userId uuid.UUID) (map[uuid.UUID]*entity.Answer, error)
}

// @Tags page
// @Security ApiKeyAuth
// @Param        id   path     string  true  "page id"
// @Success 200 {object} model.Page "ok"
// @Router /v1/page/{id} [get]
func New(pageGetter pageGetter, answersGetter answersGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var ctx = r.Context()
		_, ok := ctx.Value("userId").(uuid.UUID)

		id, err := uuid.Parse(chi.URLParam(r, "id"))
		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, err.Error())
			return
		}

		page, err := pageGetter.GetById(ctx, id)
		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, err.Error())
			return
		}

		if !ok {
			render.JSON(w, r, model.NewPage(page, map[uuid.UUID]*entity.Answer{}))
			return
		}
		panic("")
		/*
			sectionsIds := make([]uuid.UUID, 0, len(page.Sections))
			for _, section := range page.Sections {
				sectionsIds = append(sectionsIds, section.Id)
			}

			answers, err := answersGetter.GetBySections(ctx, sectionsIds, userId)
			if err != nil {
				render.Status(r, http.StatusBadRequest)
				render.JSON(w, r, err.Error())
				return
			}

			dto := model.NewPage(page, answers)
			render.Status(r, http.StatusOK)
			render.JSON(w, r, dto)

		*/
	}
}
