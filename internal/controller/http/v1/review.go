package v1

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/internal/usecase"
	"github.com/radium-rtf/radium-backend/pkg/mapper"
	"net/http"
)

type reviewRoutes struct {
	uc     usecase.ReviewUseCase
	mapper mapper.Review
}

func newReviewRoutes(useCase usecase.ReviewUseCase, signingString string) *chi.Mux {
	routes := reviewRoutes{uc: useCase}

	r := chi.NewRouter()
	r.Use(authRequired(signingString))
	r.Post("/answer", handler(routes.create).HTTP)

	return r
}

// @Tags review
// @Security ApiKeyAuth
// @Accept json
// @Param request body entity.AnswerReviewPost true "score - от 0 до 1"
// @Success 201 {object} entity.AnswerReview "created"
// @Router /review/answer [post]
func (r reviewRoutes) create(w http.ResponseWriter, request *http.Request) *appError {
	var post entity.AnswerReviewPost
	err := json.NewDecoder(request.Body).Decode(&post)
	if err != nil {
		return newAppError(err, http.StatusBadRequest)
	}
	review := r.mapper.PostToReview(post)

	review, err = r.uc.Create(request.Context(), review)
	if err != nil {
		return newAppError(err, http.StatusBadRequest)
	}

	render.Status(request, http.StatusCreated)
	render.JSON(w, request, review)
	return nil
}
