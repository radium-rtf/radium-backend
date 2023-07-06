package main

import (
	"github.com/radium-rtf/radium-backend/config"
	"github.com/radium-rtf/radium-backend/pkg/postgres/gen"
	"log"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}
	err = gen.Gen(cfg.URL)
	if err != nil {
		log.Fatal(err)
	}
}
