package main

import (
	"http-skeleton-go-1.24/src/infrastructure"
	"http-skeleton-go-1.24/src/user-interface/config"
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

	repo := &repository.MySQLRecordRepository{DB: db}
	service := &application.RecordService{Repo: repo}
	handler := &httpadapter.Handler{RecordService: service}

}
