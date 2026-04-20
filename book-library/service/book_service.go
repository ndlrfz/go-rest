package service

import (
	"book-library/model"
	"book-library/repository"
	"errors"
)

type BookService interface {
	GetByID(id uint) (*model.Book, error)
}

type bookService struct {
	repo repository.BookRepository
}

func NewBookService(repo repository.BookRepository) BookService {
	return &bookService{repo: repo}
}

func (s *bookService) GetByID(id uint) (*model.Book, error) {
	if id == 0 {
		return nil, errors.New("Invalid ID - SERVICE")
	}

	return s.repo.GetByID(id)
}
