package main

import (
	"log"
	"net/http"
)

func main() {
	cfg := loadConfig()
	connection := connectDB(cfg)

	server := newServer(connection, cfg)

	log.Printf("Server is running on port: %s", cfg.Port)
	err := http.ListenAndServe(":"+cfg.Port, server.Router)
	if err != nil {
		log.Fatalf("Server failed: %v", err)
	}

}
