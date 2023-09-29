package postgres

import (
	"database/sql"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
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
