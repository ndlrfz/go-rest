package repository

import (
	"book-library/model"
	"database/sql"

	"github.com/lib/pq"
)

type BookRepository interface {
	GetByID(id uint) (*model.Book, error)
}

type bookRepo struct {
	db *sql.DB
}

func NewBookRepository(db *sql.DB) BookRepository {
	return &bookRepo{db: db}
}

func (r *bookRepo) GetByID(id uint) (*model.Book, error) {

	var book model.Book
	query := "SELECT id, title, author, genre FROM books where id = $1"
	err := r.db.QueryRow(query, id).Scan(
		&book.ID, &book.Title, &book.Author, pq.Array(&book.Genre),
	)
	if err != nil {
		return nil, err
	}

	return &book, nil
}
