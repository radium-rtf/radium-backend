package postgres

import (
	"github.com/uptrace/bun"
	"time"
)

func maxOpenConnsDB(max int) bun.DBOption {
	return func(db *bun.DB) {
		db.SetMaxOpenConns(max)
	}
}

func maxIdleConnsDB(max int) bun.DBOption {
	return func(db *bun.DB) {
		db.SetMaxIdleConns(max)
	}
}

func connMaxLifetimeDB(max time.Duration) bun.DBOption {
	return func(db *bun.DB) {
		db.SetConnMaxLifetime(max)
	}
}

func connMaxIdleTimeDB(max time.Duration) bun.DBOption {
	return func(db *bun.DB) {
		db.SetConnMaxIdleTime(max)
	}
}
