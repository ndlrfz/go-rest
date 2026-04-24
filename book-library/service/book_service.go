package service

import (
	"book-library/model"
	"book-library/repository"
	"context"
	"errors"
)

type BookService interface {
	GetByID(ctx context.Context, id uint) (*model.Book, error)
	GetBooks(ctx context.Context) ([]model.Book, error)
	CreateBook(ctx context.Context, book *model.Book) error
}

type bookService struct {
	repo repository.BookRepository
}

func NewBookService(repo repository.BookRepository) BookService {
	return &bookService{repo: repo}
}

func (s *bookService) GetByID(ctx context.Context, id uint) (*model.Book, error) {
	if id <= 0 {
		return nil, errors.New("Invalid ID")
	}

	book, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return book, nil
}

func (s *bookService) GetBooks(ctx context.Context) ([]model.Book, error) {
	books, err := s.repo.GetBooks(ctx)
	if err != nil {
		return nil, err
	}

	return books, nil
}

func (s *bookService) CreateBook(ctx context.Context, book *model.Book) error {
	err := s.repo.CreateBook(ctx, book)
	if err != nil {
		return err
	}

	return nil
}

// package service
//
// import (
// 	"book-library/model"
// 	"book-library/repository"
// 	"errors"
// )
//
// type BookService interface {
// 	GetByID(id uint) (*model.Book, error)
// }
//
// type bookService struct {
// 	repo repository.BookRepository
// }
//
// func NewBookService(repo repository.BookRepository) BookService {
// 	return &bookService{repo: repo}
// }
//
// func (s *bookService) GetByID(id uint) (*model.Book, error) {
// 	if id == 0 {
// 		return nil, errors.New("Invalid ID - SERVICE")
// 	}
//
// 	return s.repo.GetByID(id)
// }
