package repository

import (
	"book-library/model"
	"context"
	"database/sql"
	"errors"
	"github.com/lib/pq"
	"log"
)

type BookRepository interface {
	GetByID(ctx context.Context, id uint) (*model.Book, error)
}

type bookRepository struct {
	db *sql.DB
}

func NewBookRepository(db *sql.DB) BookRepository {
	return &bookRepository{db: db}
}

func (r *bookRepository) GetByID(ctx context.Context, id uint) (*model.Book, error) {
	query := "SELECT id, title, author, genre FROM books WHERE id = $1"

	// Tambahkan log ini:
	log.Printf("DEBUG: Executing query: %s with ID: %d", query, id)

	var book model.Book
	var genre pq.StringArray // Gunakan tipe khusus dari lib/pq
	err := r.db.QueryRowContext(ctx, query, id).Scan(&book.ID, &book.Title, &book.Author, &genre)

	book.Genre = []string(genre)

	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("DEBUG: No rows found for ID: %d", id)
			return nil, errors.New("Book Not Found")
		}
		log.Printf("DEBUG: Error scanning: %v", err) // Cek apakah ada error saat scan kolom
		return nil, err
	}

	return &book, nil
}

// package repository
//
// import (
// 	"book-library/model"
// 	"database/sql"
//
// 	"github.com/lib/pq"
// )
//
// type BookRepository interface {
// 	GetByID(id uint) (*model.Book, error)
// }
//
// type bookRepo struct {
// 	db *sql.DB
// }
//
// func NewBookRepository(db *sql.DB) BookRepository {
// 	return &bookRepo{db: db}
// }
//
// func (r *bookRepo) GetByID(id uint) (*model.Book, error) {
//
// 	var book model.Book
// 	query := "SELECT id, title, author, genre FROM books where id = $1"
// 	err := r.db.QueryRow(query, id).Scan(
// 		&book.ID, &book.Title, &book.Author, pq.Array(&book.Genre),
// 	)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	return &book, nil
// }
