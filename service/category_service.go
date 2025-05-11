package service

import (
    "bookstore/model"
    "bookstore/repository"
)

type CategoryService interface {
    GetAll() ([]model.Category, error)
    Create(name string) error
    GetByID(id int) (*model.Category, error)
    Delete(id int) error
    GetBooksByCategory(id int) ([]model.Book, error)
}

type categoryService struct {
    repo *repository.CategoryRepository
}

func NewCategoryService(r *repository.CategoryRepository) CategoryService {
    return &categoryService{repo: r}
}

func (s *categoryService) GetAll() ([]model.Category, error) {
    return s.repo.GetAll()
}

func (s *categoryService) Create(name string) error {
    return s.repo.Create(name)
}

func (s *categoryService) GetByID(id int) (*model.Category, error) {
    return s.repo.GetByID(id)
}

func (s *categoryService) Delete(id int) error {
    return s.repo.Delete(id)
}

func (s *categoryService) GetBooksByCategory(id int) ([]model.Book, error) {
    return s.repo.GetBooksByCategory(id)
}