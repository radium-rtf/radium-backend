package auth

import (
	"context"
	"github.com/radium-rtf/radium-backend/pkg/auth"
	"net/http"
	"strings"
)

func UserId(manager auth.TokenManager) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			tokenHeader := strings.Split(request.Header.Get("Authorization"), " ")
			var userId any
			userId, err := manager.ExtractUserId(tokenHeader)
			if err != nil {
				userId = ""
			}
			ctx := context.WithValue(request.Context(), "userId", userId)
			next.ServeHTTP(writer, request.WithContext(ctx))
		})
	}
}
