package service

import (
	"book-library/dto"
	"book-library/model"
	"book-library/repository"
	"context"
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

type BookService interface {
	GetByID(ctx context.Context, id uint) (*model.Book, error)
	GetBooks(ctx context.Context) ([]model.Book, error)
	CreateBook(ctx context.Context, book *dto.BookReq) error
	DeleteBook(ctx context.Context, id uint) error
	UpdateBook(ctx context.Context, bookReq *dto.BookReq, id uint) error
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

func (s *bookService) CreateBook(ctx context.Context, book *dto.BookReq) error {

	isExists, err := s.repo.IsBookExists(ctx, book.Title)
	if err != nil {
		return err
	}

	if isExists {
		return fmt.Errorf("Book already exists.")
	}

	validate := validator.New()
	if err := validate.Struct(book); err != nil {
		return fmt.Errorf("Invalid input: %w", err)
	}

	err = s.repo.CreateBook(ctx, book)
	if err != nil {
		return err
	}

	return nil
}

func (s *bookService) UpdateBook(ctx context.Context, bookReq *dto.BookReq, id uint) error {
	isExists, err := s.repo.IsBookExists(ctx, bookReq.Title)
	if err != nil {
		return err
	}

	if isExists {
		return fmt.Errorf("Book already exists.")
	}

	validate := validator.New()
	if err := validate.Struct(bookReq); err != nil {
		return fmt.Errorf("Invalid input: %w", err)
	}

	err = s.repo.UpdateBook(ctx, bookReq, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *bookService) DeleteBook(ctx context.Context, id uint) error {

	err := s.repo.DeleteBook(ctx, id)
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
