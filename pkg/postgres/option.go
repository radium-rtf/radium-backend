package postgres

import "time"

type Option func(*Postgres)

func MaxOpenConns(max int) Option {
	return func(c *Postgres) {
		c.maxOpenConns = max
	}
}

func MaxIdleConns(max int) Option {
	return func(c *Postgres) {
		c.maxIdleConns = max
	}
}

func ConnMaxLifetime(max time.Duration) Option {
	return func(c *Postgres) {
		c.connMaxLifetime = max
	}
}

func ConnMaxIdleTime(max time.Duration) Option {
	return func(c *Postgres) {
		c.connMaxIdleTime = max
	}
}
