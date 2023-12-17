package postgres

import (
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/extra/bundebug"
)

type Postgres struct {
	DB *bun.DB
}

func New(url string, opts ...Option) (*Postgres, error) {
	sqldb, err := open(url)
	if err != nil {
		return nil, err
	}

	options := newOptions(opts...)

	db := bun.NewDB(sqldb, pgdialect.New(),
		maxIdleConnsDB(options.maxIdleConns),
		maxOpenConnsDB(options.maxOpenConns),
		connMaxIdleTimeDB(options.connMaxIdleTime),
		connMaxLifetimeDB(options.connMaxLifetime),
	)
	db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(false)))

	initDB(db)
	err = migrate(db)
	if err != nil {
		return nil, err
	}

	pg := &Postgres{
		DB: db,
	}

	return pg, nil
}

func (p Postgres) Close() error {
	return p.DB.Close()
}
