package postgres

import (
	"database/sql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"time"
)

func open(url string) (*gorm.DB, *sql.DB, error) {
	var (
		err    error
		gormDb *gorm.DB
	)

	for i := 0; i < 20; i++ {
		gormDb, err = gorm.Open(postgres.Open(url), &gorm.Config{})
		if err == nil {
			break
		}

		log.Println(err)
		time.Sleep(time.Second)
	}

	if err != nil {
		return nil, nil, err
	}

	sqlDb, err := gormDb.DB()
	if err != nil {
		return nil, nil, err
	}

	return gormDb, sqlDb, err
}