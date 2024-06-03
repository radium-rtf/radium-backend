package handlers

import (
	"github.com/go-chi/chi/v5"
	"github.com/radium-rtf/radium-backend/internal/wave/httpserver/handlers/chats"
	"github.com/radium-rtf/radium-backend/internal/wave/httpserver/handlers/dialogue"
	"github.com/radium-rtf/radium-backend/internal/wave/httpserver/handlers/group"
	"github.com/radium-rtf/radium-backend/internal/wave/httpserver/handlers/message"
	"github.com/radium-rtf/radium-backend/internal/wave/httpserver/handlers/user"
	"github.com/radium-rtf/radium-backend/internal/wave/usecase"
)

func routes(h chi.Router, useCases usecase.UseCases) {
	message.New(h, useCases)
	user.New(h, useCases)
	dialogue.New(h, useCases)
	group.New(h, useCases)
	chats.New(h, useCases)
}
