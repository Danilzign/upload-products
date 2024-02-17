package service

import (
	"github.com/danilzign/todo-app"
	"github.com/danilzign/todo-app/pkg/repository"
)

type Product interface {
	Create(list todo.Product) (int, error)
	GetAll() ([]todo.Product, error)
	GetById(productId int) (todo.Product, error)
	DeleteProduct(listId int) error
	UpdateProduct(listId int, input todo.UpdateProductInput) error
}

type Service struct {
	Product
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Product: NewProductService(repos.Product),
	}
}
