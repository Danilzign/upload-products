package repository

import (
	"github.com/danilzign/todo-app"
	"github.com/jmoiron/sqlx"
)

type Product interface {
	Create(list todo.Product) (int, error)
	GetAll() ([]todo.Product, error)
	GetById(productId int) (todo.Product, error)
	DeleteProduct(listId int) error
	UpdateProduct(listId int, input todo.UpdateProductInput) error
}

type Repository struct {
	Product
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Product: NewProductPostgres(db),
	}
}
