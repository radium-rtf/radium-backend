package app

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/radium-rtf/radium-backend/config"
	_ "github.com/radium-rtf/radium-backend/docs"
	"github.com/radium-rtf/radium-backend/internal/httpserver/handlers"
	"github.com/radium-rtf/radium-backend/internal/lib/session"
	"github.com/radium-rtf/radium-backend/internal/usecase"
	pg "github.com/radium-rtf/radium-backend/internal/usecase/repo/postgres"
	"github.com/radium-rtf/radium-backend/pkg/auth"
	"github.com/radium-rtf/radium-backend/pkg/filestorage"
	"github.com/radium-rtf/radium-backend/pkg/hash"
	"github.com/radium-rtf/radium-backend/pkg/httpserver"
	"github.com/radium-rtf/radium-backend/pkg/postgres"
)

func Run(cfg *config.Config) {
	db, err := postgres.New(cfg.PG.URL,
		postgres.MaxOpenConns(cfg.PG.MaxOpenConns),
		postgres.ConnMaxIdleTime(cfg.PG.ConnMaxIdleTime),
		postgres.MaxIdleConns(cfg.PG.MaxIdleConns),
		postgres.ConnMaxLifetime(cfg.PG.ConnMaxLifetime),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	storage, err := filestorage.New(cfg.Storage)
	if err != nil {
		log.Fatal(err)
	}

	r := chi.NewRouter()
	r.Use(cors.AllowAll().Handler)
	r.Get("/swagger/*", httpSwagger.Handler(httpSwagger.URL("/swagger/doc.json")))

	repositories := pg.NewRepositories(db.Q)
	tokenManager := auth.NewManager(cfg.Auth.SigningKey)
	passwordHasher := hash.NewSHA1Hasher(cfg.Auth.PasswordSalt)
	dependencies := usecase.Dependencies{
		Repos:          repositories,
		TokenManager:   tokenManager,
		Storage:        storage,
		PasswordHasher: passwordHasher,
		Session:        session.New(tokenManager, cfg.AccessTokenTTL, cfg.RefreshTokenTTL),
	}

	useCases := usecase.NewUseCases(dependencies)
	handlers.NewRouter(r, useCases)

	http := cfg.HTTP
	httpServer := httpserver.New(r,
		httpserver.Port(http.Port),
		httpserver.MaxHeaderBytes(http.MaxHeaderBytes),
		httpserver.IdleTimeout(http.IdleTimeout),
		httpserver.WriteTimeout(http.WriteTimeout),
		httpserver.ReadTimeout(http.ReadTimeout),
	)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Println("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		log.Println(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	err = httpServer.Shutdown()
	if err != nil {
		log.Println(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
}
