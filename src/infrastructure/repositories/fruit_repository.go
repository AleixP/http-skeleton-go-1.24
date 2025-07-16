package repositories

import (
	"database/sql"
	"http-skeleton-go-1.24/src/domain/model/fruit"
)

type MysqlFruitRepository struct {
	DB *sql.DB
}

func (r *MysqlFruitRepository) Create(fruit *fruit.Fruit) (int64, error) {

	result, err := r.DB.Exec("INSERT INTO fruits (name) VALUES (?)", fruit.Name)
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
	rows, err := r.DB.Query("SELECT id, name FROM fruits")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var fruits []*fruit.Fruit
	for rows.Next() {
		var item fruit.Fruit
		if err := rows.Scan(&item.ID, &item.Name); err != nil {
			return nil, err
		}
		fruits = append(fruits, &item)
	}
	return fruits, nil
}
