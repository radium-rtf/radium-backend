package postgres

import (
	"database/sql/driver"
	"github.com/google/uuid"
)

type uuids []uuid.UUID

func (v uuids) toValuers() []driver.Valuer {
	valuers := make([]driver.Valuer, 0, len(v))
	for _, id := range v {
		valuers = append(valuers, id)
	}
	return valuers
}

type columnValue struct {
	column string
	value  any
}

type columnValues struct {
	column string
	value  any
}
