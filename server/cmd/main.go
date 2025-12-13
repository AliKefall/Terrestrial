package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	cfg := loadConfig()
	connection := connectDB(cfg)

	server := newServer(connection, cfg)

	log.Printf("Server is running on port: %s", cfg.Port)
	err := http.ListenAndServe(":"+cfg.Port, server.Router)
	if err != nil {
		log.Fatalf("Server failed: %v", err)
	}

}
