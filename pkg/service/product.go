package service

import (
	"github.com/danilzign/todo-app"
	"github.com/danilzign/todo-app/pkg/repository"
)

type TodoListService struct {
	repo repository.Product
}

func NewProductService(repo repository.Product) *TodoListService {
	return &TodoListService{repo: repo}
}

func (s *TodoListService) Create(product todo.Product) (int, error) {
	return s.repo.Create(product)
}

func (s *TodoListService) GetAll() ([]todo.Product, error) {
	return s.repo.GetAll()

}

func (s *TodoListService) GetById(productId int) (todo.Product, error) {
	return s.repo.GetById(productId)

}

func (s *TodoListService) DeleteProduct(listId int) error {
	return s.repo.DeleteProduct(listId)

}

func (s *TodoListService) UpdateProduct(productId int, input todo.UpdateProductInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.UpdateProduct(productId, input)
}
