package main

import (
	"database/sql"
	"github.com/joho/godotenv"
	"http-skeleton-go-1.24/src/domain/service"
	"http-skeleton-go-1.24/src/infrastructure"
	"http-skeleton-go-1.24/src/infrastructure/repository"
	"http-skeleton-go-1.24/src/user-interface/config"
	"log"
	"net/http"
)

func main() {
	initEnv()
	db := initDb()
	fruitService := initServices(db)
	router := config.NewRouter(fruitService)

	log.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))

}

func initServices(db *sql.DB) *service.FruitService {
	fruitRepository := repository.NewMysqlFruitRepository(db)
	fruitService := service.NewFruitService(fruitRepository)
	return fruitService
}

func initDb() *sql.DB {
	cfg, _ := config.Load()

	db, err := infrastructure.StartMySQL(cfg)
	if err != nil {
		log.Fatalf("Cannot connect to DB: %v", err)
	}

	if err := infrastructure.Migrate(db); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
	return db
}

func initEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}
}
