package infrastructure

import (
	"database/sql"
	"http-skeleton-go-1.24/src/config"
)

func StartMySQL(cfg *config.Config) (*sql.DB, error) {
	db, err := sql.Open("database", cfg.DSN())
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
