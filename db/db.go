package db

import (
	"ecom/types"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewSQLStorage(cfg postgres.Config) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(cfg.DSN), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	var All = []interface{}{
		&types.User{},
		&types.Product{},
		&types.Order{},
		&types.OrderItem{},
	}

	er := db.AutoMigrate(All...)
	if er != nil {
		log.Fatal(er)
	}

	return db, nil
}
