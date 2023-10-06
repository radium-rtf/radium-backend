package migrations

import (
	"github.com/uptrace/bun/migrate"
	"os"
)

var Migrations = migrate.NewMigrations()

func init() {
	if err := Migrations.Discover(os.DirFS("./")); err != nil {
		panic(err)
	}
}
