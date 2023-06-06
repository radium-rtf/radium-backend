package v1

import (
	"encoding/json"
	"net/http"

	"github.com/dranikpg/dto-mapper"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/pkg/errors"

	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/internal/usecase"
)

type courseRoutes struct {
	uc usecase.CourseUseCase
}

func newCourseRoutes(h chi.Router, useCase usecase.CourseUseCase, signingKey string) {
	routes := courseRoutes{uc: useCase}

	h.Route("/course", func(r chi.Router) {
		r.Get("/", handler(routes.getCourses).HTTP)
		r.Get("/{courseId}", handler(routes.getCourse).HTTP)

		r.Group(func(r chi.Router) {
			r.Use(authRequired(signingKey))
			r.Post("/", handler(routes.postCourse).HTTP)
		})
	})

	h.Group(func(r chi.Router) {
		// r.Use(authRequired(signingKey))
		// r.Post("/link/course", handler(routes.postLink).HTTP)
		// r.Post("/collaborator", handler(routes.postCollaborator).HTTP)
		// r.Post("/join/course/{courseId}", handler(routes.join).HTTP)
	})
}

// @Tags course
// @Security ApiKeyAuth
// @Accept       json
// @Param request body entity.CourseRequest true "Данные о курсе"
// @Success      201   {object} entity.CourseDto       "created"
// @Router       /course [post]
func (r courseRoutes) postCourse(w http.ResponseWriter, request *http.Request) *appError {
	courseRequest := &entity.CourseRequest{}
	err := json.NewDecoder(request.Body).Decode(&courseRequest)
	if err != nil {
		return newAppError(errors.New(err.Error()), http.StatusBadRequest)
	}

	course, err := r.uc.CreateCourse(request.Context(), *courseRequest)
	if err != nil {
		return newAppError(err, http.StatusBadRequest)
	}

	c := entity.CourseDto{}
	dto.Map(&c, course)

	render.Status(request, http.StatusCreated)
	render.JSON(w, request, c)

	return nil
}

// @Tags course
// @Security ApiKeyAuth
// @Success      200   {object} entity.Course        "ok"
// @Router       /course [get]
func (r courseRoutes) getCourses(w http.ResponseWriter, request *http.Request) *appError {
	courses, err := r.uc.GetCourses(request.Context())
	if err != nil {
		return newAppError(err, http.StatusBadRequest)
	}

	c := []entity.CourseDto{}
	dto.Map(&c, courses)

	render.Status(request, http.StatusOK)
	render.JSON(w, request, c)
	return nil
}

// @Tags course
// @Param        courseId   path     string  true  "course id"
// @Success      200   {object} entity.CourseDto  "ok"
// @Router       /course/{courseId} [get]
func (r courseRoutes) getCourse(w http.ResponseWriter, request *http.Request) *appError {
	courseId := chi.URLParam(request, "courseId")
	course, err := r.uc.GetCourseById(request.Context(), courseId)
	if err != nil {
		return newAppError(err, http.StatusBadRequest)
	}

	c := entity.CourseDto{}
	dto.Map(&c, course)
	render.Status(request, http.StatusOK)
	render.JSON(w, request, c)
	return nil
}

// // @Tags course
// // @Security ApiKeyAuth
// // @Param        courseId   path      integer  true  "course id"
// // @Success      201   {object} entity.Course "created"
// // @Router       /join/course/{courseId} [post]
// func (r courseRoutes) join(w http.ResponseWriter, request *http.Request) *appError {
// 	userId := request.Context().Value("userId").(string)
// 	courseId := chi.URLParam(request, "courseId")
// 	courses, err := r.uc.Join(request.Context(), userId, courseId)
// 	if err != nil {
// 		return newAppError(err, http.StatusBadRequest)
// 	}
// 	render.Status(request, http.StatusCreated)
// 	render.JSON(w, request, courses)
// 	return nil
// }
