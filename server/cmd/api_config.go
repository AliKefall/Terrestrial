package main

import (
	"log"
	"os"
)

type Config struct {
	DatabaseURL   string
	DatabaseToken string
	TokenSecret   string
	Port          string
}

func loadConfig() *Config {
	cfg := &Config{
		DatabaseURL:   os.Getenv("TURSO_URL"),
		DatabaseToken: os.Getenv("TURSO_AUTH_TOKEN"),
		Port:          getEnv("PORT", "8080"),
	}
	if cfg.DatabaseToken == "" {
		log.Fatal("Database connection link environment variable is missing.")

	}
	if cfg.DatabaseToken == "" {
		log.Fatal("Database authentication token is missing.")
	}
	return cfg

}

func getEnv(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return fallback

}
