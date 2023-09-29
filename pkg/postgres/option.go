package postgres

import (
	"time"
)

const (
	defaultMaxOpenConns    = 3
	defaultMaxIdleConns    = 2
	defaultConnMaxIdleTime = time.Minute * 1
	defaultConnMaxLifetime = time.Minute * 1
)

type options struct {
	maxOpenConns    int
	maxIdleConns    int
	connMaxIdleTime time.Duration
	connMaxLifetime time.Duration
}

func newOptions(opts ...Option) *options {
	options := &options{
		maxIdleConns:    defaultMaxIdleConns,
		maxOpenConns:    defaultMaxOpenConns,
		connMaxLifetime: defaultConnMaxLifetime,
		connMaxIdleTime: defaultConnMaxIdleTime,
	}

	for _, op := range opts {
		op(options)
	}

	return options
}

type Option func(*options)

func MaxOpenConns(max int) Option {
	return func(c *options) {
		c.maxOpenConns = max
	}
}

func MaxIdleConns(max int) Option {
	return func(c *options) {
		c.maxIdleConns = max
	}
}

func ConnMaxLifetime(max time.Duration) Option {
	return func(c *options) {
		c.connMaxLifetime = max
	}
}

func ConnMaxIdleTime(max time.Duration) Option {
	return func(c *options) {
		c.connMaxIdleTime = max
	}
}
