package main

import (
	"fmt"
	"go-restapi/app"
	"go-restapi/controller"
	"go-restapi/middleware"
	"go-restapi/repository"
	"go-restapi/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	bookRepository := repository.NewBookRepository()
	bookService := service.NewBookServiceImpl(bookRepository, db, validate)
	bookController := controller.NewBookController(bookService)
	router := app.NewRouter(bookController)
	fmt.Println(db)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	fmt.Println("server running")
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
