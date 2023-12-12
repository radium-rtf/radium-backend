package app

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/radium-rtf/radium-backend/config"
	_ "github.com/radium-rtf/radium-backend/docs"
	"github.com/radium-rtf/radium-backend/pkg/filestorage"
)

func Run(cfg *config.Config) {
	db, err := openDB(cfg.PG)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	storage := filestorage.New(cfg.Storage)

	dependencies := newDependencies(storage, cfg, db)

	httpServer := startHttpServer(cfg.HTTP, dependencies)

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
