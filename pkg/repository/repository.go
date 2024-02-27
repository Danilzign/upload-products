package repository

import (
	test "test"

	"github.com/jmoiron/sqlx"
)

type Product interface {
	Create(list test.Product) (string, error)
	GetAll(limit string, page string) ([]test.Product, error)
	GetById(productId string) (test.Product, error)
	DeleteProduct(productId string) (string, error)
	UpdateProduct(productId string, input test.UpdateProductInput) error
}

type Repository struct {
	Product
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Product: NewProductPostgres(db),
	}
}
