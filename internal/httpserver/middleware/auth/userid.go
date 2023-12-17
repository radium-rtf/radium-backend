package auth

import (
	"context"
	"github.com/go-chi/httplog/v2"
	"github.com/radium-rtf/radium-backend/pkg/auth"
	"log/slog"
	"net/http"
	"strings"
)

func UserId(manager auth.TokenManager) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			tokenHeader := strings.Split(request.Header.Get("Authorization"), " ")

			userId, err := manager.ExtractUserId(tokenHeader)
			if err != nil {
				next.ServeHTTP(writer, request)
				return
			}

			httplog.LogEntrySetField(request.Context(), "user", slog.StringValue(userId.String()))
			ctx := context.WithValue(request.Context(), "userId", userId)
			next.ServeHTTP(writer, request.WithContext(ctx))
		})
	}
}
