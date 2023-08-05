package postgres

import (
	"database/sql"
	"github.com/radium-rtf/radium-backend/pkg/postgres/db"
	"github.com/radium-rtf/radium-backend/pkg/postgres/pggen"
	"time"
)

const (
	defaultMaxOpenConns    = 3
	defaultMaxIdleConns    = 2
	defaultConnMaxIdleTime = time.Minute * 1
	defaultConnMaxLifetime = time.Minute * 1
)

type Postgres struct {
	Q   *db.Query
	sql *sql.DB

	maxOpenConns    int
	maxIdleConns    int
	connMaxIdleTime time.Duration
	connMaxLifetime time.Duration
}

func New(url string, opts ...Option) (*Postgres, error) {
	gormDb, sqlDb, err := open(url)

	if err != nil {
		return nil, err
	}

	err = pggen.Gen()
	if err != nil {
		return nil, err
	}

	Q := db.Use(gormDb)

	err = migrate(gormDb)

	if err != nil {
		return nil, err
	}

	pg := &Postgres{
		Q:               Q,
		sql:             sqlDb,
		maxOpenConns:    defaultMaxOpenConns,
		maxIdleConns:    defaultMaxIdleConns,
		connMaxLifetime: defaultConnMaxLifetime,
		connMaxIdleTime: defaultConnMaxIdleTime,
	}

	for _, opt := range opts {
		opt(pg)
	}

	sqlDb.SetMaxOpenConns(pg.maxOpenConns)
	sqlDb.SetMaxIdleConns(pg.maxIdleConns)
	sqlDb.SetConnMaxIdleTime(pg.connMaxIdleTime)
	sqlDb.SetConnMaxLifetime(pg.connMaxLifetime)

	return &Postgres{Q: Q, sql: sqlDb}, err
}

func (p Postgres) Close() error {
	return p.sql.Close()
}
