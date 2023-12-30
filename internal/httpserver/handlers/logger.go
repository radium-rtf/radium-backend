package handlers

import (
	"fmt"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httplog/v2"
	"log/slog"
	"net/http"
	"strings"
)

func newLogger() *httplog.Logger {
	logger := httplog.NewLogger("radium-http", httplog.Options{
		JSON:             true,
		LogLevel:         slog.LevelInfo,
		Concise:          true,
		RequestHeaders:   true,
		MessageFieldName: "message",
	})
	return logger
}

func newHandlerLogger(log *httplog.Logger) func(h http.Handler) http.Handler {
	paths := []string{"/auth/", "/password", "/swagger/"}
	logfn := httplog.Handler(log)

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			for _, path := range paths {
				if strings.Contains(request.URL.Path, path) {
					var f middleware.LogFormatter = &requestLogger{*log.Logger, log.Options}

					entry := f.NewLogEntry(request)
					next.ServeHTTP(writer, middleware.WithLogEntry(request, entry))
					return
				}
			}
			logfn(next).ServeHTTP(writer, request)
		})
	}
}

type requestLogger struct {
	Logger  slog.Logger
	Options httplog.Options
}

func (l *requestLogger) NewLogEntry(r *http.Request) middleware.LogEntry {
	entry := &httplog.RequestLoggerEntry{}
	msg := fmt.Sprintf("Request: %s %s", r.Method, r.URL.Path)

	if !l.Options.Concise {
		entry.Logger.Info(msg)
	}
	return entry
}
