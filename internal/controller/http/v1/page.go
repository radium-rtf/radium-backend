package v1

import (
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/pkg/mapper"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/internal/usecase"
)

type pageRoutes struct {
	uc     usecase.PageUseCase
	mapper mapper.Page
}

func newPageRoutes(h chi.Router, useCase usecase.PageUseCase, signingString string) {
	routes := pageRoutes{uc: useCase}

	h.Route("/page", func(r chi.Router) {
		// r.Get("/{slideId}", handler(routes.getSlideSections).HTTP)
		// r.Get/("/{courseId}/{moduleNameEng}", handler(routes.getSlides).HTTP)
		r.Get("/{id}", handler(routes.getById).HTTP)
		r.Group(func(r chi.Router) {
			r.Use(authToken(signingString))
			r.Get("/{id}", handler(routes.getById).HTTP)
		})

		r.Group(func(r chi.Router) {
			r.Use(authRequired(signingString))
			r.Post("/", handler(routes.postSlide).HTTP)
		})
	})
}

// @Tags page
// @Security ApiKeyAuth
// @Param       request body entity.PageRequest true "создание слайда"
// @Success      201   {object} entity.PageDto "ok"
// @Router      /page [post]
func (r pageRoutes) postSlide(w http.ResponseWriter, request *http.Request) *appError {
	var pageRequest entity.PageRequest
	if err := render.DecodeJSON(request.Body, &pageRequest); err != nil {
		return newAppError(err, http.StatusBadRequest)
	}
	pageDto, err := r.uc.CreatePage(request.Context(), pageRequest)
	if err != nil {
		return newAppError(err, http.StatusBadRequest)
	}

	render.Status(request, http.StatusCreated)
	render.JSON(w, request, pageDto)
	return nil
}

// @Tags page
// @Security ApiKeyAuth
// @Param        id   path     string  true  "page id"
// @Success 200 {object} entity.PageDto "ok"
// @Router /page/{id} [get]
func (r pageRoutes) getById(w http.ResponseWriter, request *http.Request) *appError {
	id := chi.URLParam(request, "id")
	uid, err := uuid.Parse(id)
	if err != nil {
		return newAppError(err, http.StatusBadRequest)
	}

	userId, ok := request.Context().Value("userId").(uuid.UUID)
	var pageDto *entity.PageDto
	if !ok {
		pageDto, err = r.uc.GetByID(request.Context(), uid, nil)
	} else {
		pageDto, err = r.uc.GetByID(request.Context(), uid, &userId)
	}

	if err != nil {
		return newAppError(err, http.StatusBadRequest)
	}

	render.Status(request, http.StatusOK)
	render.JSON(w, request, pageDto)

	return nil
}

// // @Tags slide
// // @Security ApiKeyAuth
// // @Param        courseId   path      integer  true  "course id"
// // @Param        moduleNameEng   path     string  true  "moduleName"
// // @Success      201   {object} entity.ModuleSlides "ok"
// // @Router      /slide/{courseId}/{moduleNameEng} [get]
// func (r slideRoutes) getSlides(w http.ResponseWriter, request *http.Request) *appError {
// 	courseId, err := strconv.Atoi(chi.URLParam(request, "courseId"))
// 	if err != nil {
// 		return newAppError(err, http.StatusBadRequest)
// 	}
// 	moduleNameEng := chi.URLParam(request, "moduleNameEng")
// 	slideRequest := entity.SlidesRequest{ModuleNameEng: moduleNameEng, CourseId: uint(courseId)}
// 	slide, err := r.uc.GetSlides(request.Context(), slideRequest)
// 	if err != nil {
// 		return newAppError(err, http.StatusBadRequest)
// 	}
// 	render.Status(request, http.StatusOK)
// 	render.JSON(w, request, slide)
// 	return nil
// }

// // @Tags slide
// // @Security ApiKeyAuth
// // @Param        slideId   path      integer  true  "slide id"
// // @Success      201   {object} entity.SlideSections "ok"
// // @Router      /slide/{slideId} [get]
// func (r slideRoutes) getSlideSections(w http.ResponseWriter, request *http.Request) *appError {
// 	slideId, err := strconv.Atoi(chi.URLParam(request, "slideId"))
// 	if err != nil {
// 		return newAppError(err, http.StatusBadRequest)
// 	}
// 	slideSectionsRequest := entity.SlideSectionsRequest{SlideId: uint(slideId)}
// 	slide, err := r.uc.GetSlideSections(request.Context(), slideSectionsRequest)
// 	if err != nil {
// 		return newAppError(err, http.StatusBadRequest)
// 	}
// 	render.Status(request, http.StatusOK)
// 	render.JSON(w, request, slide)
// 	return nil
// }
