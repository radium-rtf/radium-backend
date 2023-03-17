package app

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	httpSwagger "github.com/swaggo/http-swagger"
	"log"

	"github.com/radium-rtf/radium-backend/config"
	_ "github.com/radium-rtf/radium-backend/docs"
	v1 "github.com/radium-rtf/radium-backend/internal/controller/http/v1"
	"github.com/radium-rtf/radium-backend/pkg/filestorage"
	"github.com/radium-rtf/radium-backend/pkg/httpserver"
	"github.com/radium-rtf/radium-backend/pkg/postgres"
)

func Run(cfg *config.Config) {
	migrator(cfg)
	pg, err := postgres.New(cfg.PG.URL, postgres.MaxPoolSize(cfg.PG.PoolMax))
	if err != nil {
		log.Fatal(err)
	}
	defer pg.Close()
	r := chi.NewRouter()
	r.Use(cors.AllowAll().Handler)
	r.Get("/swagger/*", httpSwagger.Handler(httpSwagger.URL("/swagger/doc.json")))
	storage := filestorage.New(cfg.Storage)
	v1.NewRouter(r, pg, storage, cfg)

	httpServer := httpserver.New(r, httpserver.Port(cfg.HTTP.Port))
	log.Fatal(httpServer.ListenAndServe())
}
