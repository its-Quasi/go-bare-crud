package store

import (
	"bare-crud/internal/models"
	"database/sql"
)

type Store interface {
	GetAll() ([]models.Book, error)
	GetById(id int) (models.Book, error)
	Create(book models.Book) (models.Book, error)
	// Update(id int) (models.Book, error)
	// Delete(id int) (models.Book, error)
}

type store struct {
	db *sql.DB
}

func New(db *sql.DB) Store {
	return &store{db: db}
}

func (s *store) GetAll() ([]models.Book, error) {
	q := `SELECT id, title, description FROM books`
	rows, err := s.db.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var books []models.Book
	for rows.Next() {
		var b models.Book
		if err := rows.Scan(&b.Id, &b.Title, &b.Description); err != nil {
			return nil, err
		}
		books = append(books, b)
	}
	return books, nil
}

func (s *store) GetById(id int) (models.Book, error) {
	q := `SELECT * FROM books WHERE id = ?`
	var book models.Book
	err := s.db.QueryRow(q, id).Scan(&book.Id, book.Description, book.Title)
	if err != nil {
		return models.Book{}, err
	}
	return book, nil
}

func (s *store) Create(bookData models.Book) (models.Book, error) {
	q := `INSERT INTO TABLE books (title, description) VALUES (?,?)`
	res, err := s.db.Exec(q, bookData.Description, bookData.Title)
	if err != nil {
		return models.Book{}, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return models.Book{}, err
	}
	bookData.Id = uint(id)
	return bookData, nil
}
