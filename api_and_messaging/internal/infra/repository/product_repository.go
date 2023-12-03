package repository

import (
	"database/sql"

	"github.com/GuiCintra27/go/api_and_messaging/internal/entity"
)

type ProductRepositoryMySQL struct {
	DB *sql.DB
}

func NewProductRepositoryMySQL(db *sql.DB) *ProductRepositoryMySQL {
	return &ProductRepositoryMySQL{DB: db}
}

func (repository *ProductRepositoryMySQL) Create(product *entity.Product) error {
	_, err := repository.DB.Exec("INSERT INTO products (id, name, price) VALUES (?, ?, ?)", product.Id, product.Name, product.Price)
	
	if err != nil {
		return err
	}

	return nil
}

func (repository *ProductRepositoryMySQL) FindAll() ([]*entity.Product, error) {
	rows, err := repository.DB.Query("SELECT id, name, price FROM products")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*entity.Product

	for rows.Next() {
		var product entity.Product
		if err := rows.Scan(&product.Id, &product.Name, &product.Price); err != nil {
			return nil, err
		}
		products = append(products, &product)
	}

	return products, nil
}