package auth

import (
	"context"
	"github.com/go-chi/httplog/v2"
	"github.com/radium-rtf/radium-backend/internal/lib/auth"
	"log/slog"
	"net/http"
	"strings"
)

func Required(manager auth.TokenManager) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			tokenHeader := strings.Split(request.Header.Get("Authorization"), " ")

			userId, err := manager.ExtractUserId(tokenHeader)
			if err != nil {
				writer.WriteHeader(http.StatusUnauthorized)
				writer.Write([]byte(err.Error()))
				return
			}

			httplog.LogEntrySetField(request.Context(), "user", slog.StringValue(userId.String()))

			ctx := context.WithValue(request.Context(), "userId", userId)
			next.ServeHTTP(writer, request.WithContext(ctx))
		})
	}
}
