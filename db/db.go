package db

import (
	"ecom/types"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var All = []interface{}{
	&types.User{},
	&types.Product{},
	&types.Order{},
	&types.OrderItem{},
}

func NewSQLStorage(cfg postgres.Config) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(cfg.DSN), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	er := MigrateAll(db)
	if er != nil {
		log.Fatal(er)
	}

	return db, nil
}

func MigrateAll(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.AutoMigrate(All...); err != nil {
			return err
		}
		return nil
	})
}
