package create

import (
	"context"
	"encoding/json"
	"github.com/go-chi/render"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/internal/lib/answer/verdict"
	"github.com/radium-rtf/radium-backend/internal/model"
	"github.com/radium-rtf/radium-backend/pkg/validator"
	"net/http"
)

type creator interface {
	Create(ctx context.Context, section *entity.Section) (*entity.Section, error)
}

// @Tags section
// @Security ApiKeyAuth
// @Accept json
// @Param request body Section true "Информация для раздела"
// @Success 201 {object} model.Section "created"
// @Router /v1/section [post]
func New(creator creator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			create Section
			ctx    = r.Context()
		)

		err := json.NewDecoder(r.Body).Decode(&create)
		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, err.Error())
			return
		}

		err = validator.Struct(create)
		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, err.Error())
			return
		}

		section := create.toSection()
		section, err = creator.Create(ctx, section)
		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, err.Error())
			return
		}

		dto := model.NewSection(section, verdict.EMPTY, 0, "", nil)
		render.Status(r, http.StatusCreated)
		render.JSON(w, r, dto)
	}
}
