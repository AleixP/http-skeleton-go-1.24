package main

import (
	"database/sql"
	"github.com/joho/godotenv"
	"http-skeleton-go-1.24/src/domain/services"
	"http-skeleton-go-1.24/src/infrastructure"
	"http-skeleton-go-1.24/src/infrastructure/repositories"
	"http-skeleton-go-1.24/src/user-interface/config"
	"log"
	"net/http"
)

func main() {
	initEnv()
	db := initDb()
	fruitService := initServices(db)
	router := config.NewRouter(fruitService)
	
	log.Println("Server running at http://localhost:8081")
	log.Fatal(http.ListenAndServe(":8081", router))

}

func initServices(db *sql.DB) *services.FruitService {
	fruitRepository := repositories.NewMysqlFruitRepository(db)
	fruitService := services.NewFruitService(fruitRepository)
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
