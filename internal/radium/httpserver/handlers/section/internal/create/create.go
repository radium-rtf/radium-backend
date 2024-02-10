package create

import (
	"context"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/radium/entity"
	"github.com/radium-rtf/radium-backend/internal/radium/model"
	"github.com/radium-rtf/radium-backend/pkg/decode"
	"github.com/radium-rtf/radium-backend/pkg/resp"
	"net/http"
)

type creator interface {
	Create(ctx context.Context, section *entity.Section, editorId uuid.UUID) (*entity.Section, error)
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
			userId = r.Context().Value("userId").(uuid.UUID)
		)

		err := decode.Json(r.Body, &create)
		if err != nil {
			resp.Error(r, w, err)
			return
		}

		section, err := create.ToSection()
		if err != nil {
			resp.Error(r, w, err)
			return
		}

		section, err = creator.Create(ctx, section, userId)
		if err != nil {
			resp.Error(r, w, err)
			return
		}

		dto := model.NewSection(section, nil, int(section.MaxAttempts.Int16))
		render.Status(r, http.StatusCreated)
		render.JSON(w, r, dto)
	}
}
