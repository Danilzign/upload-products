package service

import (
	"fmt"
	"os"
	test "test"
	"test/pkg/repository"
)

type ProductService struct {
	repo repository.Product
}

func NewProductService(repo repository.Product) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) Create(product test.Product) (string, error) {
	return s.repo.Create(product)
}

func (s *ProductService) GetAll(limit string, page string) ([]test.Product, error) {
	return s.repo.GetAll(limit, page)

}

func (s *ProductService) GetById(productId string) (test.Product, error) {
	return s.repo.GetById(productId)

}

func (s *ProductService) DeleteProduct(productId string) (string, error) {
	imageName, err := s.repo.DeleteProduct(productId)
	if err != nil {
		return "", err
	}

	directory := fmt.Sprintf("dev/TaskTest/image/product/default/%s", imageName)
	os.Remove(directory)

	return s.repo.DeleteProduct(productId)
}

func (s *ProductService) UpdateProduct(productId string, input test.UpdateProductInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.UpdateProduct(productId, input)
}
