package main

import (
	"github.com/radium-rtf/radium-backend/config"
	"github.com/radium-rtf/radium-backend/pkg/postgres"
	"log"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}
	_, err = postgres.New(cfg.URL)
	if err != nil {
		log.Fatal(err)
	}
}
