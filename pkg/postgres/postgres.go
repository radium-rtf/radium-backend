package postgres

import (
	"database/sql"
	"time"
)

const (
	defaultMaxOpenConns    = 3
	defaultMaxIdleConns    = 2
	defaultConnMaxIdleTime = time.Minute * 1
	defaultConnMaxLifetime = time.Minute * 1
)

type Postgres struct {
	sql *sql.DB

	maxOpenConns    int
	maxIdleConns    int
	connMaxIdleTime time.Duration
	connMaxLifetime time.Duration
}

func New(url string, opts ...Option) (*Postgres, error) {
	panic("not implemented")
}

func (p Postgres) Close() error {
	return p.sql.Close()
}
