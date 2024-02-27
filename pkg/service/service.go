package service

import (
	test "test"
	"test/pkg/repository"
)

type Product interface {
	Create(list test.Product) (string, error)
	GetAll(limit string, page string) ([]test.Product, error)
	GetById(productId string) (test.Product, error)
	DeleteProduct(productId string) (string, error)
	UpdateProduct(productId string, input test.UpdateProductInput) error
}

type Service struct {
	Product
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Product: NewProductService(repos.Product),
	}
}
