package main

import (
	"ecom/cmd/api"
	"ecom/config"
	"ecom/db"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
)

func main() {

	dataDB, err := db.NewSQLStorage(postgres.Config{
		DSN: fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", config.Envs.DBHost, config.Envs.DBPort, config.Envs.DBUser, config.Envs.DBPassword, config.Envs.DBName),
	})

	if err != nil {
		log.Fatal(err)
	}
	server := api.NewServer(":8080", dataDB)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
