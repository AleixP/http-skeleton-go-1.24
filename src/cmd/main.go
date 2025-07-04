package main

import (
	"http-skeleton-go-1.24/src/config"
	"http-skeleton-go-1.24/src/infrastructure"
	"log"
)

func main() {
	cfg, _ := config.Load()

	db, err := infrastructure.StartMySQL(cfg)
	if err != nil {
		log.Fatalf("Cannot connect to DB: %v", err)
	}

	if err := infrastructure.Migrate(db); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
}
