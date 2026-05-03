package repository

import (
	"context"
	"database/sql"
	"go-restapi/model/domain"
)

type BookRepository interface {
	Save(ctx context.Context, tx *sql.Tx, book domain.Book) domain.Book
	Update(ctx context.Context, tx *sql.Tx, book domain.Book) domain.Book
	Delete(ctx context.Context, tx *sql.Tx, book domain.Book)
	FindById(ctx context.Context, tx *sql.Tx, bookId string) (domain.Book, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Book
}
