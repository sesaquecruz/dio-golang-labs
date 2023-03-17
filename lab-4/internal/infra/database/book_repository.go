package database

import (
	"database/sql"
	"lab-4/internal/entity"

	_ "github.com/mattn/go-sqlite3"
)

type BookRepository struct {
	Db *sql.DB
}

func NewBookRepository(db *sql.DB) *BookRepository {
	return &BookRepository{Db: db}
}

func (r *BookRepository) FindAll() ([]entity.Book, error) {
	rows, err := r.Db.Query("SELECT id, title, author, category, year FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	books := []entity.Book{}
	for rows.Next() {
		var book entity.Book
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Category, &book.Year)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (r *BookRepository) FindById(id int64) (*entity.Book, error) {
	var book entity.Book
	row := r.Db.QueryRow("SELECT id, title, author, category, year FROM books WHERE id=?", id)
	err := row.Scan(&book.ID, &book.Title, &book.Author, &book.Category, &book.Year)
	if err != nil {
		return nil, err
	}
	return &book, nil
}

func (r *BookRepository) Save(book *entity.Book) (int64, error) {
	stmt, err := r.Db.Prepare("INSERT INTO books (title, author, category, year) VALUES (?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}
	res, err := stmt.Exec(book.Title, book.Author, book.Category, book.Year)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func (r *BookRepository) Update(id int64, book *entity.Book) (int64, error) {
	stmt, err := r.Db.Prepare("UPDATE books SET title=?, author=?, category=?, year=? WHERE id=?")
	if err != nil {
		return 0, err
	}
	res, err := stmt.Exec(book.Title, book.Author, book.Category, book.Year, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func (r *BookRepository) Delete(id int64) (int64, error) {
	stmt, err := r.Db.Prepare("DELETE FROM books WHERE id=?")
	if err != nil {
		return 0, nil
	}
	res, err := stmt.Exec(id)
	if err != nil {
		return 0, nil
	}
	return res.RowsAffected()
}
