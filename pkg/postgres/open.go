package postgres

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"time"
)

func open(url string) (*gorm.DB, error) {
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

	return gormDb, err
}
