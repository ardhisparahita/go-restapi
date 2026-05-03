package service

import (
	"context"
	"go-restapi/model/web"
)

type BookService interface {
	Create(ctx context.Context, r web.BookCreateRequest) web.BookResponse
	Update(ctx context.Context, r web.BookUpdateRequest) web.BookResponse
	Delete(ctx context.Context, bookId string)
	FindById(ctx context.Context, bookId string) web.BookResponse
	FindAll(ctx context.Context) []web.BookResponse
}
