package v1

import (
	"encoding/json"
	"github.com/google/uuid"
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

func newCourseRoutes(useCase usecase.CourseUseCase, signingKey string) *chi.Mux {
	routes := courseRoutes{uc: useCase}
	r := chi.NewRouter()

	r.Get("/", handler(routes.getCourses).HTTP)
	r.Get("/{courseId}", handler(routes.getCourse).HTTP)
	r.Get("/slug/{slug}", handler(routes.getCourseBySlug).HTTP)

	r.Group(func(r chi.Router) {
		r.Use(authRequired(signingKey))
		r.Post("/", handler(routes.postCourse).HTTP)
		r.Patch("/join/{courseId}", handler(routes.join).HTTP)
		r.Delete("/{id}", handler(routes.delete).HTTP)
	})

	r.Group(func(r chi.Router) {
		// r.Use(authRequired(signingKey))
		// r.Post("/link/course", handler(routes.postLink).HTTP)
		// r.Post("/join/course/{courseId}", handler(routes.join).HTTP)
	})

	return r
}

// @Tags course
// @Security ApiKeyAuth
// @Accept       json
// @Param request body entity.CoursePost true "Данные о курсе"
// @Success      201   {object} entity.CourseDto       "created"
// @Router       /course [post]
func (r courseRoutes) postCourse(w http.ResponseWriter, request *http.Request) *appError {
	courseRequest := &entity.CoursePost{}
	err := json.NewDecoder(request.Body).Decode(&courseRequest)
	if err != nil {
		return newAppError(errors.New(err.Error()), http.StatusBadRequest)
	}

	course := entity.NewCourse(*courseRequest)
	course, err = r.uc.CreateCourse(request.Context(), course)
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
// @Success      200   {object} entity.CourseDto        "ok"
// @Router       /course [get]
func (r courseRoutes) getCourses(w http.ResponseWriter, request *http.Request) *appError {
	courses, err := r.uc.GetCourses(request.Context())
	if err != nil {
		return newAppError(err, http.StatusBadRequest)
	}

	c := make([]entity.CourseDto, 0, len(courses))
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
	courseId, err := uuid.Parse(chi.URLParam(request, "courseId"))
	if err != nil {
		return newAppError(err, http.StatusBadRequest)
	}
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

// @Tags course
// @Param        slug   path     string  true  "course slug"
// @Success      200   {object} entity.CourseDto  "ok"
// @Router       /course/slug/{slug} [get]
func (r courseRoutes) getCourseBySlug(w http.ResponseWriter, request *http.Request) *appError {
	slug := chi.URLParam(request, "slug")
	course, err := r.uc.GetCourseBySlug(request.Context(), slug)
	if err != nil {
		return newAppError(err, http.StatusBadRequest)
	}

	c := entity.CourseDto{}
	dto.Map(&c, course)
	render.Status(request, http.StatusOK)
	render.JSON(w, request, c)
	return nil
}

// @Tags course
// @Security ApiKeyAuth
// @Param        courseId   path      string  true  "course id"
// @Success      201   {object} entity.CourseDto "created"
// @Router       /course/join/{courseId} [patch]
func (r courseRoutes) join(w http.ResponseWriter, request *http.Request) *appError {
	userId := request.Context().Value("userId").(uuid.UUID)
	courseId, err := uuid.Parse(chi.URLParam(request, "courseId"))
	if err != nil {
		return newAppError(err, http.StatusCreated)
	}
	courses, err := r.uc.Join(request.Context(), userId, courseId)
	if err != nil {
		return newAppError(err, http.StatusBadRequest)
	}
	render.Status(request, http.StatusCreated)
	render.JSON(w, request, courses)
	return nil
}

// @Tags course
// @Security ApiKeyAuth
// @Accept json
// @Param        id   path      string  true  "id"
// @Param   is_soft     query     boolean    false  "по умолчанию soft"
// @Success 200
// @Router /course/{id} [delete]
func (r courseRoutes) delete(w http.ResponseWriter, request *http.Request) *appError {
	destroy, err := newDestroy(request)
	if err != nil {
		return newAppError(err, http.StatusBadRequest)
	}

	err = r.uc.Delete(request.Context(), destroy)
	if err != nil {
		return newAppError(err, http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusOK)
	return nil
}
