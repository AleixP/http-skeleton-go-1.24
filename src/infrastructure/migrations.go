package infrastructure

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func ensureMigrationsTableIsUp(db *sql.DB) error {
	const query = `
	CREATE TABLE IF NOT EXISTS migrations (
		id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(255) NOT NULL UNIQUE,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`
	_, err := db.Exec(query)
	return err
}

func getMigrations(db *sql.DB) (map[string]bool, error) {
	applied := make(map[string]bool)
	rows, err := db.Query("SELECT name FROM migrations")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var name string
	for rows.Next() {
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}
		applied[name] = true
	}
	return applied, nil
}

func Migrate(db *sql.DB) error {
	if err := ensureMigrationsTableIsUp(db); err != nil {
		return err
	}

	applied, err := getMigrations(db)
	if err != nil {
		return err
	}

	files, err := filepath.Glob("database/migrations/*.sql")
	if err != nil {
		return err
	}
	sort.Strings(files)

	for _, file := range files {
		name := filepath.Base(file)
		if applied[name] {
			continue // already exists
		}
		sqlBytes, err := os.ReadFile(file)
		if err != nil {
			return err
		}

		sqlStmt := strings.TrimSpace(string(sqlBytes))
		if sqlStmt == "" { // empty file
			continue
		}
		log.Printf("Applying migration: %s", name)
		if _, err := db.Exec(sqlStmt); err != nil {
			return err
		}

		if _, err := db.Exec("INSERT INTO migrations (name) VALUES (?)", name); err != nil {
			return err
		}
	}
	log.Println("Migrations applied successfully")
	return nil
}
