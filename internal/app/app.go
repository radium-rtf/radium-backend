package app

import (
	"fmt"
	"github.com/radium-rtf/radium-backend/internal/httpserver/handlers"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/radium-rtf/radium-backend/config"
	_ "github.com/radium-rtf/radium-backend/docs"
	"github.com/radium-rtf/radium-backend/pkg/filestorage"
	"github.com/radium-rtf/radium-backend/pkg/httpserver"
	"github.com/radium-rtf/radium-backend/pkg/postgres"
)

func Run(cfg *config.Config) {
	pg := cfg.PG
	db, err := postgres.New(pg.URL,
		postgres.MaxOpenConns(pg.MaxOpenConns),
		postgres.ConnMaxIdleTime(pg.ConnMaxIdleTime),
		postgres.MaxIdleConns(pg.MaxIdleConns),
		postgres.ConnMaxLifetime(pg.ConnMaxLifetime),
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
	handlers.NewRouter(r, db.Q, storage, cfg)

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
