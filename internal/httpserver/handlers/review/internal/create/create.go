package create

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"net/http"

	"github.com/go-chi/render"
	"github.com/radium-rtf/radium-backend/internal/entity"
)

type creator interface {
	Create(ctx context.Context, review *entity.Review) (*entity.Review, error)
}

// @Tags review
// @Security ApiKeyAuth
// @Accept json
// @Param request body Review true "score - от 0 до 1"
// @Success 201 {object} entity.Review "created"
// @Router /v1/review [post]
func NewReview(creator creator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			request Review
			ctx     = r.Context()
			userId  = ctx.Value("userId").(uuid.UUID)
		)

		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, err.Error())
			return
		}

		review := request.toReview(userId)
		review, err = creator.Create(ctx, review)
		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, err.Error())
			return
		}

		render.Status(r, http.StatusCreated)
		render.JSON(w, r, review)
	}
}
