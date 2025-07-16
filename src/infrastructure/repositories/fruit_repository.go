package repositories

import (
	"database/sql"
	"http-skeleton-go-1.24/src/domain/model/fruit"
)

type MysqlFruitRepository struct {
	DB *sql.DB
}

func NewMysqlFruitRepository(db *sql.DB) *MysqlFruitRepository {
	return &MysqlFruitRepository{DB: db}
}

func (r *MysqlFruitRepository) Create(fruit *fruit.Fruit) (int64, error) {
	query := `INSERT INTO fruits (name) VALUES (?)`
	result, err := r.DB.Exec(query, fruit.Name)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *MysqlFruitRepository) List() ([]*fruit.Fruit, error) {
	rows, err := r.DB.Query("SELECT id, name, color FROM fruits")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var fruits []*fruit.Fruit
	for rows.Next() {
		var item fruit.Fruit
		if err := rows.Scan(&item.ID, &item.Name, &item.Color); err != nil {
			return nil, err
		}
		fruits = append(fruits, &item)
	}
	return fruits, nil
}

func (r *MysqlFruitRepository) FindById(id string) (*fruit.Fruit, error) {
	var item fruit.Fruit
	query := `SELECT id, name, color FROM fruits WHERE id = ?`

	err := r.DB.QueryRow(query, id).Scan(
		&item.ID,
		&item.Name,
		&item.Color)
	if err != nil {
		return nil, err
	}
	return &item, nil
}
