package v1

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/internal/usecase"
	"net/http"
	"strconv"
)

type moduleRoutes struct {
	uc usecase.ModuleUseCase
}

func newModuleRoutes(course chi.Router, useCase usecase.ModuleUseCase, signingKey string) {
	routes := moduleRoutes{uc: useCase}
	course.Get("/{courseId}", handler(routes.getModules).HTTP)

	course.Group(func(r chi.Router) {
		r.Use(authRequired(signingKey))
		r.Post("/{courseId}", handler(routes.postModule).HTTP)
	})
}

// @Tags module
// @Security ApiKeyAuth
// @Param       request body entity.ModuleRequest true "moduleRequest"
// @Param        courseId   path     integer  true  "course id"
// @Success      201   {object} entity.ModuleDto       "created"
// @Router       /course/{courseId} [post]
func (r moduleRoutes) postModule(w http.ResponseWriter, request *http.Request) *appError {
	var module entity.ModuleRequest
	courseId := chi.URLParam(request, "courseId")
	id, err := strconv.Atoi(courseId)
	if err != nil {
		return newAppError(err, http.StatusBadRequest)
	}
	if err = render.DecodeJSON(request.Body, &module); err != nil {
		return newAppError(err, http.StatusBadRequest)
	}
	moduleDto, err := r.uc.CreateModule(request.Context(), id, module)
	if err != nil {
		return newAppError(err, http.StatusBadRequest)
	}
	render.Status(request, http.StatusCreated)
	render.JSON(w, request, moduleDto)
	return nil
}

// @Tags module
// @Security ApiKeyAuth
// @Param        courseId   path      integer  true  "course id"
// @Success      200   {object} entity.CourseModules "ok"
// @Router       /course/{courseId} [get]
func (r moduleRoutes) getModules(w http.ResponseWriter, request *http.Request) *appError {
	courseId := chi.URLParam(request, "courseId")
	id, err := strconv.Atoi(courseId)
	if err != nil {
		return newAppError(err, http.StatusBadRequest)
	}

	courseModules, err := r.uc.GetCourseModules(request.Context(), id)
	if err != nil {
		return newAppError(err, http.StatusBadRequest)
	}
	render.JSON(w, request, courseModules)
	return nil
}
