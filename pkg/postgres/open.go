package postgres

import (
	"database/sql"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"time"
)

func open(url string) (*sql.DB, error) {
	config, err := pgx.ParseConfig(url)
	if err != nil {
		return nil, err
	}

	sqldb := stdlib.OpenDB(*config)
	if err = sqldb.Ping(); err == nil {
		return sqldb, nil
	}

	for i := 0; i < 20; i++ {
		time.Sleep(time.Second * 20)
		if err = sqldb.Ping(); err == nil {
			return sqldb, nil
		}
	}

	return nil, err
}

func OpenDB(url string) (*bun.DB, error) {
	sqldb, err := open(url)
	if err != nil {
		return nil, err
	}
	db := bun.NewDB(sqldb, pgdialect.New())
	return db, nil
}
