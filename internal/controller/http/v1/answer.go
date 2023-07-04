package v1

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/internal/usecase"
	"github.com/radium-rtf/radium-backend/pkg/mapper"
	"net/http"
)

type answerRoutes struct {
	uc     usecase.AnswerUseCase
	mapper mapper.Answer
}

func newAnswerRoutes(useCase usecase.AnswerUseCase, signingString string) *chi.Mux {
	routes := answerRoutes{uc: useCase}
	r := chi.NewRouter()
	r.Group(func(r chi.Router) {
		r.Use(authRequired(signingString))
		r.Post("/", handler(routes.createAnswer).HTTP)
	})
	return r
}

// @Tags answer
// @Security ApiKeyAuth
// @Param       request body entity.AnswerPost true "ответ"
// @Success      201   {object} entity.SectionDto "ok"
// @Router      /answer [post]
func (r answerRoutes) createAnswer(w http.ResponseWriter, request *http.Request) *appError {
	post := &entity.AnswerPost{}
	if err := render.DecodeJSON(request.Body, post); err != nil {
		return newAppError(err, http.StatusBadRequest)
	}
	userId := request.Context().Value("userId").(uuid.UUID)
	answer := entity.NewPostToAnswer(post, userId)
	answer, err := r.uc.Answer(request.Context(), answer)
	if err != nil {
		return newAppError(err, http.StatusBadRequest)
	}

	dto := r.mapper.Answer(answer)

	render.Status(request, http.StatusCreated)
	render.JSON(w, request, dto)
	return nil
}
