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

type teacherRoutes struct {
	uc usecase.TeacherUseCase
	mapper.Teacher
}

func newTeacherRoutes(useCase usecase.TeacherUseCase, signingKey string) chi.Router {
	routes := teacherRoutes{uc: useCase}
	r := chi.NewRouter()

	r.Use(authRequired(signingKey))
	r.Post("/", handler(routes.post).HTTP)
	r.Get("/courses", handler(routes.courses).HTTP)
	return r
}

// @Tags teacher
// @Security ApiKeyAuth
// @Param       request body entity.TeacherPost true " "
// @Success      201   {string}  string        "created"
// @Router       /teacher [post]
func (r teacherRoutes) post(w http.ResponseWriter, request *http.Request) *appError {
	var post entity.TeacherPost
	err := render.DecodeJSON(request.Body, &post)
	if err != nil {
		return newAppError(err, http.StatusBadRequest)
	}
	teacher := r.PostToTeacher(post)
	teacher, err = r.uc.Create(request.Context(), teacher)
	if err != nil {
		return newAppError(err, http.StatusBadRequest)
	}
	render.Status(request, http.StatusCreated)
	render.JSON(w, request, teacher)
	return nil
}

// @Tags teacher
// @Security ApiKeyAuth
// @Success      200   {string}  entity.TeacherCourse        " "
// @Router       /teacher/courses [get]
func (r teacherRoutes) courses(w http.ResponseWriter, request *http.Request) *appError {
	userId := request.Context().Value("userId").(uuid.UUID)
	teacher, err := r.uc.GetByUserId(request.Context(), userId)
	if err != nil {
		return newAppError(err, http.StatusBadRequest)
	}
	courses := r.ToCourses(teacher)
	render.Status(request, http.StatusOK)
	render.JSON(w, request, courses)
	return nil
}
