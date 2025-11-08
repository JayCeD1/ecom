package config

import (
	"os"
	"strconv"
)

type Config struct {
	DBUser                 string
	DBPassword             string
	DBHost                 string
	DBPort                 string
	DBName                 string
	JWTExpirationInSeconds int64
	JWTSecret              string
}

var Envs = initConfig()

func initConfig() Config {
	return Config{
		DBUser:                 getEnv("DB_USER", "ecom"),
		DBPassword:             getEnv("DB_PASSWORD", "addidas1"),
		DBHost:                 getEnv("DB_HOST", "localhost"),
		DBPort:                 getEnv("DB_PORT", "5432"),
		DBName:                 getEnv("DB_NAME", "ecom"),
		JWTExpirationInSeconds: getEnvAsInt("JWT_EXP", 3600),
		JWTSecret:              getEnv("JWT_SECRET", "This_composite_literal_allocates_given_values."),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getEnvAsInt(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.ParseInt(value, 10, 64)
		if err == nil {
			return i
		}
	}
	return fallback
}
