package main

import (
	"database/sql"
	"fmt"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
	"log"
)

func connectDB(cfg *Config) *sql.DB {
	dsn := fmt.Sprintf("%s?authToken=%s", cfg.DatabaseURL, cfg.DatabaseToken)
	db, err := sql.Open("libsql", dsn)
	if err != nil {
		log.Fatalf("Failed to open Turso database: %v", err)
	}

	db.SetMaxIdleConns(3)
	db.SetMaxOpenConns(3)

	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to connect to Turso: %v", err)
	}

	log.Println("Connected to Turso successfully")
	return db
}
