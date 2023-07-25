package main

import (
	"github.com/radium-rtf/radium-backend/pkg/postgres/pggen"
	"log"
)

func main() {
	err := pggen.Gen()
	if err != nil {
		log.Fatal(err)
	}
}
