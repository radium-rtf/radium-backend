package app

import (
	"fmt"
	"github.com/radium-rtf/radium-backend/pkg/httpserver"
	"github.com/radium-rtf/radium-backend/pkg/postgres"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/radium-rtf/radium-backend/config"
	_ "github.com/radium-rtf/radium-backend/docs"
	"github.com/radium-rtf/radium-backend/pkg/filestorage"
)

type App struct {
	httpServer *httpserver.Server
}

func NewApp(cfg *config.Config, db *postgres.Postgres) App {
	storage := filestorage.New(cfg.Storage)
	dependencies := newDependencies(storage, cfg, db)
	httpServer := newHttpServer(cfg.Radium.HTTP, dependencies)

	return App{httpServer: httpServer}
}

func (app App) Run() error {
	app.httpServer.Start()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Println("radium - app - Run - signal: " + s.String())
	case err := <-app.httpServer.Notify():
		log.Println(fmt.Errorf("wave - app - Run - httpServer.Notify: %w", err))
	}

	err := app.shutdown()
	return err
}

func (app App) shutdown() error {
	err := app.httpServer.Shutdown()
	if err != nil {
		log.Println(fmt.Errorf("wave - app - Run - httpServer.Shutdown: %w", err))
	}
	return err
}
