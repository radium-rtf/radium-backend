package main

import (
	"fmt"
	bunmigrate "github.com/uptrace/bun/migrate"
	"github.com/urfave/cli/v2"
	"strings"
)

type migrator struct {
	migrator *bunmigrate.Migrator
}

func newMigrator(bunMigrator *bunmigrate.Migrator) *migrator {
	return &migrator{migrator: bunMigrator}
}

func (m *migrator) migrate(c *cli.Context) error {
	if err := m.migrator.Lock(c.Context); err != nil {
		return err
	}
	defer m.migrator.Unlock(c.Context)

	group, err := m.migrator.Migrate(c.Context)
	if err != nil {
		return err
	}
	if group.IsZero() {
		fmt.Printf("there are no newMigrator migrations to run (database is up to date)\n")
		return nil
	}
	fmt.Printf("migrated to %s\n", group)
	return nil
}

func (m *migrator) init(c *cli.Context) error {
	return m.migrator.Init(c.Context)
}

func (m *migrator) rollback(c *cli.Context) error {
	if err := m.migrator.Lock(c.Context); err != nil {
		return err
	}
	defer m.migrator.Unlock(c.Context)

	group, err := m.migrator.Rollback(c.Context)
	if err != nil {
		return err
	}
	if group.IsZero() {
		fmt.Printf("there are no groups to roll back\n")
		return nil
	}
	fmt.Printf("rolled back %s\n", group)
	return nil
}

func (m *migrator) lock(c *cli.Context) error {
	return m.migrator.Lock(c.Context)
}

func (m *migrator) createGO(c *cli.Context) error {
	name := strings.Join(c.Args().Slice(), "_")
	mf, err := m.migrator.CreateGoMigration(c.Context, name)
	if err != nil {
		return err
	}
	fmt.Printf("created migration %s (%s)\n", mf.Name, mf.Path)
	return nil
}

func (m *migrator) createSQL(c *cli.Context) error {
	name := strings.Join(c.Args().Slice(), "_")
	files, err := m.migrator.CreateSQLMigrations(c.Context, name)
	if err != nil {
		return err
	}

	for _, mf := range files {
		fmt.Printf("created migration %s (%s)\n", mf.Name, mf.Path)
	}

	return nil
}

func (m *migrator) status(c *cli.Context) error {
	ms, err := m.migrator.MigrationsWithStatus(c.Context)
	if err != nil {
		return err
	}
	fmt.Printf("migrations: %s\n", ms)
	fmt.Printf("unapplied migrations: %s\n", ms.Unapplied())
	fmt.Printf("last migration group: %s\n", ms.LastGroup())
	return nil
}

func (m *migrator) markApplied(c *cli.Context) error {
	group, err := m.migrator.Migrate(c.Context, bunmigrate.WithNopMigration())
	if err != nil {
		return err
	}
	if group.IsZero() {
		fmt.Printf("there are no newMigrator migrations to mark as applied\n")
		return nil
	}
	fmt.Printf("marked as applied %s\n", group)
	return nil
}

func (m *migrator) unlock(c *cli.Context) error {
	return m.migrator.Unlock(c.Context)
}
