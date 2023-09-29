package main

import (
	"github.com/radium-rtf/radium-backend/config"
	"github.com/radium-rtf/radium-backend/migrations"
	"github.com/radium-rtf/radium-backend/pkg/postgres"
	"github.com/uptrace/bun/migrate"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalln(err)
	}

	db, err := postgres.New(cfg.URL)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	migrator := newMigrator(migrate.NewMigrator(db.DB, migrations.Migrations))
	app := &cli.App{
		Name: "bun",

		Commands: []*cli.Command{
			newDBCommand(migrator),
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func newDBCommand(migrator *migrator) *cli.Command {
	return &cli.Command{
		Name:  "db",
		Usage: "database migrations",
		Subcommands: []*cli.Command{
			{
				Name:   "init",
				Usage:  "create migration tables",
				Action: migrator.init,
			},
			{
				Name:   "migrate",
				Usage:  "migrate database",
				Action: migrator.migrate,
			},
			{
				Name:   "rollback",
				Usage:  "rollback the last migration group",
				Action: migrator.rollback,
			},
			{
				Name:   "lock",
				Usage:  "lock migrations",
				Action: migrator.lock,
			},
			{
				Name:   "unlock",
				Usage:  "unlock migrations",
				Action: migrator.unlock,
			},
			{
				Name:   "create_go",
				Usage:  "create Go migration",
				Action: migrator.createGO,
			},
			{
				Name:   "create_sql",
				Usage:  "create up and down SQL migrations",
				Action: migrator.createSQL,
			},
			{
				Name:   "status",
				Usage:  "print migrations status",
				Action: migrator.status,
			},
			{
				Name:   "mark_applied",
				Usage:  "mark migrations as applied without actually running them",
				Action: migrator.markApplied,
			},
		},
	}
}
