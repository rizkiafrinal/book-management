package service

import (
    "bookstore/model"
    "bookstore/repository"
)

type BookService interface {
    GetAll() ([]model.Book, error)
    Create(b *model.Book) error
    GetByID(id int) (*model.Book, error)
    Delete(id int) error
}

type bookService struct {
    repo *repository.BookRepository
}

func NewBookService(r *repository.BookRepository) BookService {
    return &bookService{repo: r}
}

func (s *bookService) GetAll() ([]model.Book, error) {
    return s.repo.GetAll()
}

func (s *bookService) Create(b *model.Book) error {
    return s.repo.Create(b)
}

func (s *bookService) GetByID(id int) (*model.Book, error) {
    return s.repo.GetByID(id)
}

func (s *bookService) Delete(id int) error {
    return s.repo.Delete(id)
}