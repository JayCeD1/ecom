package main

import (
	"ecom/cmd/api"
	"ecom/config"
	"ecom/db"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"gorm.io/driver/postgres"
)

func main() {

	dataDB, err := db.NewSQLStorage(postgres.Config{
		DSN: fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", config.Envs.DBHost, config.Envs.DBPort, config.Envs.DBUser, config.Envs.DBPassword, config.Envs.DBName),
	})

	if err != nil {
		log.Fatal(err)
	}

	sqlDB, err := dataDB.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := sqlDB.Close(); err != nil {
			log.Printf("failed to close database: %v", err)
		}
	}()

	server := api.NewServer(":8080", dataDB)

	// Graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		log.Println("Gracefully Shutting down...")
		_ = server.App.Shutdown()
	}()

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
