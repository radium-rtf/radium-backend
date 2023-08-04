package main

import (
	"log"

	"github.com/radium-rtf/radium-backend/pkg/postgres/pggen"
)

func main() {
	err := pggen.Gen()
	if err != nil {
		log.Fatal(err)
	}
}
