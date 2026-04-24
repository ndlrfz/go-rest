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
	GetBooks(ctx context.Context) ([]model.Book, error)
	CreateBook(ctx context.Context, book *model.Book) error
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

func (r *bookRepository) GetBooks(ctx context.Context) ([]model.Book, error) {
	var books []model.Book
	query := "SELECT title, genre, author FROM books"

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		log.Printf("DEBUG: Error row context %v", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var book model.Book
		var genre pq.StringArray

		err := rows.Scan(&book.Title, &genre, &book.Author)
		if err != nil {
			log.Printf("DEBUG: Error scan rows %v", err)
			return nil, err
		}

		book.Genre = []string(genre)

		books = append(books, book)
	}

	if err := rows.Err(); err != nil {
		log.Printf("DEBUG: Error tidak ada rows %v", err)
		return nil, err
	}

	return books, nil
}

func (r *bookRepository) CreateBook(ctx context.Context, book *model.Book) error {
	query := "INSERT INTO books (title, genre, author) VALUES ($1, $2, $3)"

	_, err := r.db.ExecContext(ctx, query, book.Title, pq.Array(book.Genre), book.Author)
	if err != nil {
		return err
	}
	// log.Printf("Data masuk: %+v", book)

	return nil
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
