package repository

import (
	"context"
	"database/sql"
	"errors"
	"go-restapi/model/domain"
)

type BookRepositoryImpl struct {
}

func NewBookRepository() BookRepository {
	return &BookRepositoryImpl{}
}

func (repository *BookRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, book domain.Book) domain.Book {
	SQL := "INSERT INTO books(id, title, author) VALUES(?, ?, ?)"
	_, err := tx.ExecContext(ctx, SQL, book.Id, book.Title, book.Author)
	if err != nil {
		panic(err)
	}
	return book
}

func (repository *BookRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, book domain.Book) domain.Book {
	var existing domain.Book

	err := tx.QueryRowContext(ctx, "SELECT title, author FROM books WHERE id = ?", book.Id).
		Scan(&existing.Title, &existing.Author)
	if err != nil {
		panic(err)
	}

	if book.Title == "" {
		book.Title = existing.Title
	}
	if book.Author == "" {
		book.Author = existing.Author
	}

	_, err = tx.ExecContext(ctx,
		"UPDATE books SET title = ?, author = ? WHERE id = ?",
		book.Title, book.Author, book.Id,
	)
	if err != nil {
		panic(err)
	}
	return book
}

func (repository *BookRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, book domain.Book) {
	SQL := "DELETE FROM books WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, book.Id)
	if err != nil {
		panic(err)
	}
}

func (repository *BookRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, bookId string) (domain.Book, error) {
	SQL := "SELECT id, title, author FROM books WHERE id=?"
	rows, err := tx.QueryContext(ctx, SQL, bookId)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	book := domain.Book{}

	if rows.Next() {
		err := rows.Scan(&book.Id, &book.Title, &book.Author)
		if err != nil {
			panic(err)
		}
		return book, nil
	} else {
		return book, errors.New("Books Not Found")
	}
}

func (repository *BookRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Book {
	SQL := "SELECT id, title, author FROM books"
	rows, err := tx.QueryContext(ctx, SQL)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var books []domain.Book

	for rows.Next() {
		book := domain.Book{}
		err := rows.Scan(&book.Id, &book.Title, &book.Author)
		if err != nil {
			panic(err)
		}
		books = append(books, book)
	}
	return books
}
