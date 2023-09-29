package postgres

import (
	"context"
	"fmt"
	"github.com/radium-rtf/radium-backend/migrations"
	"github.com/uptrace/bun"
	bunmigrate "github.com/uptrace/bun/migrate"
)

func migrate(db *bun.DB) error {
	migrator := bunmigrate.NewMigrator(db, migrations.Migrations)
	ctx := context.Background()

	err := migrator.Init(ctx)
	if err != nil {
		return err
	}

	if err = migrator.Lock(ctx); err != nil {
		return err
	}
	defer migrator.Unlock(ctx)

	group, err := migrator.Migrate(ctx)
	if err != nil {
		return err
	}
	if group.IsZero() {
		fmt.Printf("there are no new migrations to run (database is up to date)\n")
		return nil
	}

	fmt.Printf("migrated to %s\n", group)
	return nil
}
