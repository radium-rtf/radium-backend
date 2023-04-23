package v1

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/internal/usecase"
	"net/http"
	"strconv"
)

type slideRoutes struct {
	uc usecase.SlideUseCase
}

func newSlideRoutes(h chi.Router, useCase usecase.SlideUseCase, signingString string) {
	routes := slideRoutes{uc: useCase}

	h.Route("/slide", func(r chi.Router) {
		r.Get("/{courseId}/{moduleNameEng}", handler(routes.getSlides).HTTP)
		r.Group(func(r chi.Router) {
			r.Use(authRequired(signingString))
			r.Post("/", handler(routes.postSlide).HTTP)
		})
	})
}

// @Tags slide
// @Security ApiKeyAuth
// @Param       request body entity.SlideRequest true "создание слайда"
// @Success      201   {object} entity.Slide "ok"
// @Router      /slide [post]
func (r slideRoutes) postSlide(w http.ResponseWriter, request *http.Request) *appError {
	var slidePost entity.SlideRequest
	if err := render.DecodeJSON(request.Body, &slidePost); err != nil {
		return newAppError(err, http.StatusBadRequest)
	}
	slide, err := r.uc.CreateSlide(request.Context(), slidePost)
	if err != nil {
		return newAppError(err, http.StatusBadRequest)
	}
	render.Status(request, http.StatusCreated)
	render.JSON(w, request, slide)
	return nil
}

// @Tags slide
// @Security ApiKeyAuth
// @Param        courseId   path      integer  true  "course id"
// @Param        moduleNameEng   path     string  true  "moduleName"
// @Success      201   {object} entity.ModuleSlides "ok"
// @Router      /slide/{courseId}/{moduleNameEng} [get]
func (r slideRoutes) getSlides(w http.ResponseWriter, request *http.Request) *appError {
	courseId, err := strconv.Atoi(chi.URLParam(request, "courseId"))
	if err != nil {
		return newAppError(err, http.StatusBadRequest)
	}
	moduleNameEng := chi.URLParam(request, "moduleNameEng")
	slideRequest := entity.SlidesRequest{ModuleNameEng: moduleNameEng, CourseId: uint(courseId)}
	slide, err := r.uc.GetSlides(request.Context(), slideRequest)
	if err != nil {
		return newAppError(err, http.StatusBadRequest)
	}
	render.Status(request, http.StatusOK)
	render.JSON(w, request, slide)
	return nil
}
