package service

import (
	"context"
	"database/sql"
	"go-restapi/exception"
	"go-restapi/model/domain"
	"go-restapi/model/web"
	"go-restapi/repository"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type BookServiceImpl struct {
	BookRepository repository.BookRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewBookServiceImpl(bookRepository repository.BookRepository, DB *sql.DB, validate *validator.Validate) BookService {
	return &BookServiceImpl{
		BookRepository: bookRepository,
		DB:             DB,
		Validate:       validate,
	}
}

func (service *BookServiceImpl) Create(ctx context.Context, r web.BookCreateRequest) web.BookResponse {
	err := service.Validate.Struct(r)
	if err != nil {
		panic(err)
	}
	tx, err := service.DB.Begin()
	if err != nil {
		panic(err)
	}
	defer func() {
		err := recover()
		if err != nil {
			tx.Rollback()
			panic(err)
		} else {
			tx.Commit()
		}
	}()

	book := domain.Book{
		Id:     uuid.NewString(),
		Title:  r.Title,
		Author: r.Author,
	}

	book = service.BookRepository.Save(ctx, tx, book)
	return web.BookResponse{
		Id:     book.Id,
		Title:  book.Title,
		Author: book.Author,
	}
}

func (service *BookServiceImpl) FindById(ctx context.Context, bookId string) web.BookResponse {
	tx, err := service.DB.Begin()
	if err != nil {
		panic(err)
	}
	defer func() {
		err := recover()
		if err != nil {
			tx.Rollback()
			panic(err)
		} else {
			tx.Commit()
		}
	}()

	book, err := service.BookRepository.FindById(ctx, tx, bookId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return web.BookResponse{
		Id:     book.Id,
		Title:  book.Title,
		Author: book.Author,
	}

}

func (service *BookServiceImpl) FindAll(ctx context.Context) []web.BookResponse {
	tx, err := service.DB.Begin()
	if err != nil {
		panic(err)
	}
	defer func() {
		err := recover()
		if err != nil {
			tx.Rollback()
			panic(err)
		} else {
			tx.Commit()
		}
	}()

	var bookResponses []web.BookResponse

	books := service.BookRepository.FindAll(ctx, tx)

	for _, book := range books {
		bookResponses = append(bookResponses, web.BookResponse{
			Id:     book.Id,
			Title:  book.Title,
			Author: book.Author,
		})
	}
	return bookResponses
}

func (service *BookServiceImpl) Update(ctx context.Context, r web.BookUpdateRequest) web.BookResponse {
	err := service.Validate.Struct(r)
	if err != nil {
		panic(err)
	}
	tx, err := service.DB.Begin()
	if err != nil {
		panic(err)
	}
	defer func() {
		err := recover()
		if err != nil {
			tx.Rollback()
			panic(err)
		} else {
			tx.Commit()
		}
	}()

	book, err := service.BookRepository.FindById(ctx, tx, r.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	book.Title = r.Title
	book.Author = r.Author

	book = service.BookRepository.Update(ctx, tx, book)

	return web.BookResponse{
		Id:     book.Id,
		Title:  book.Title,
		Author: book.Author,
	}
}

func (service *BookServiceImpl) Delete(ctx context.Context, bookId string) {
	tx, err := service.DB.Begin()
	if err != nil {
		panic(err)
	}
	defer func() {
		err := recover()
		if err != nil {
			tx.Rollback()
			panic(err)
		} else {
			tx.Commit()
		}
	}()

	book, err := service.BookRepository.FindById(ctx, tx, bookId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.BookRepository.Delete(ctx, tx, book)
}
