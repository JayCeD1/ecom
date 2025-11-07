package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewSQLStorage(cfg postgres.Config) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(cfg.DSN), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return db, nil
}
