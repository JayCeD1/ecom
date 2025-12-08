package main

import (
	"context"
	"ecom/cmd/api"
	"ecom/config"
	"ecom/db"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

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
		<-c // 1. Wait for signal (Ctrl+C or K8s SIGTERM)
		log.Println("Gracefully Shutting down...")

		// Give Fiber time to finish requests
		// 2. Create a timeout context (10-second max)
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel() // 3. Cleanup context when done

		// 4. Tell Fiber to shut down with timeout
		if err = server.App.ShutdownWithContext(ctx); err != nil {
			log.Printf("Server forced to shutdown: %v", err)
		}
	}()

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
