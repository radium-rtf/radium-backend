package app

import (
	"errors"
	"fmt"
	"github.com/radium-rtf/radium-backend/pkg/httpserver"
	"github.com/radium-rtf/radium-backend/pkg/postgres"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/radium-rtf/radium-backend/config"
	_ "github.com/radium-rtf/radium-backend/docs/radium"
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

	var err error
	select {
	case s := <-interrupt:
		err = errors.New("radium - app - Run - signal: " + s.String())
	case err = <-app.httpServer.Notify():
		err = fmt.Errorf("wave - app - Run - httpServer.Notify: %w", err)
	}

	return err
}

func (app App) Shutdown() error {
	err := app.httpServer.Shutdown()
	if err != nil {
		log.Println(fmt.Errorf("wave - app - Run - httpServer.Shutdown: %w", err))
	}
	return err
}
