package handlers

import (
	"github.com/go-chi/chi/v5"
	"github.com/radium-rtf/radium-backend/internal/radium/httpserver/handlers/account"
	"github.com/radium-rtf/radium-backend/internal/radium/httpserver/handlers/answer"
	"github.com/radium-rtf/radium-backend/internal/radium/httpserver/handlers/auth"
	"github.com/radium-rtf/radium-backend/internal/radium/httpserver/handlers/author"
	"github.com/radium-rtf/radium-backend/internal/radium/httpserver/handlers/course"
	"github.com/radium-rtf/radium-backend/internal/radium/httpserver/handlers/file"
	"github.com/radium-rtf/radium-backend/internal/radium/httpserver/handlers/group"
	"github.com/radium-rtf/radium-backend/internal/radium/httpserver/handlers/module"
	notification "github.com/radium-rtf/radium-backend/internal/radium/httpserver/handlers/notification"
	"github.com/radium-rtf/radium-backend/internal/radium/httpserver/handlers/page"
	"github.com/radium-rtf/radium-backend/internal/radium/httpserver/handlers/review"
	"github.com/radium-rtf/radium-backend/internal/radium/httpserver/handlers/role"
	"github.com/radium-rtf/radium-backend/internal/radium/httpserver/handlers/section"
	"github.com/radium-rtf/radium-backend/internal/radium/httpserver/handlers/teacher"
	"github.com/radium-rtf/radium-backend/internal/radium/httpserver/handlers/user"
	"github.com/radium-rtf/radium-backend/internal/radium/usecase"
)

func routes(h chi.Router, useCases usecase.UseCases) {
	course.New(h, useCases)
	module.New(h, useCases)
	page.New(h, useCases)
	section.New(h, useCases)

	group.New(h, useCases)
	teacher.New(h, useCases)
	review.New(h, useCases)

	answer.New(h, useCases)

	auth.New(h, useCases)
	account.New(h, useCases)
	user.New(h, useCases)
	role.New(h, useCases)

	file.New(h, useCases)

	author.New(h, useCases)

	notification.New(h, useCases)
}
