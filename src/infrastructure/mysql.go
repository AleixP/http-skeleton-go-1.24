package infrastructure

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"http-skeleton-go-1.24/src/config"
)

func StartMySQL(cfg *config.Config) (*sql.DB, error) {
	db, err := sql.Open("mysql", cfg.DSN())
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
