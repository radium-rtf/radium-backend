package app

import (
	"errors"
	"github.com/golang-migrate/migrate/v4"
	"github.com/radium-rtf/radium-backend/config"
	"log"
	"time"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

const (
	_defaultAttempts = 30
	_defaultTimeout  = time.Second * 2
)

func migrator(cfg *config.Config) {
	databaseURL := cfg.PG.URL
	var (
		attempts = _defaultAttempts
		err      error
		m        *migrate.Migrate
	)
	log.Print(databaseURL)
	for attempts > 0 {
		m, err = migrate.New("file://migrations", databaseURL)
		if err == nil {
			break
		}

		log.Printf("Migrate: postgres is trying to connect, attempts left: %d, err: %s", attempts, err.Error())
		time.Sleep(_defaultTimeout)
		attempts--
	}

	if err != nil {
		log.Fatalf("Migrate: postgres connect error: %s", err)
	}

	err = m.Up()
	defer m.Close()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatalf("Migrate: up error: %s", err)
	}

	if errors.Is(err, migrate.ErrNoChange) {
		log.Printf("Migrate: no change")
		return
	}

	log.Printf("Migrate: up success")
}
