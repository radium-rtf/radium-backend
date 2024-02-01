package handlers

import (
	"github.com/go-chi/httplog/v2"
	"log/slog"
	"net/http"
	"strings"
)

func newLogger() *httplog.Logger {
	logger := httplog.NewLogger("wave-http", httplog.Options{
		JSON:             true,
		LogLevel:         slog.LevelInfo,
		Concise:          true,
		RequestHeaders:   true,
		MessageFieldName: "message",
	})
	return logger
}

func newHandlerLogger(log *httplog.Logger) func(h http.Handler) http.Handler {
	paths := []string{"/swagger/"}
	logfn := httplog.Handler(log)

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			for _, path := range paths {
				if strings.Contains(request.URL.Path, path) {
					next.ServeHTTP(writer, request)
					return
				}
			}
			logfn(next).ServeHTTP(writer, request)
		})
	}
}
