package app

import (
	"go-restapi/controller"
	"go-restapi/exception"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(bookController controller.BookController) *httprouter.Router {
	router := httprouter.New()

	router.POST("/api/books", bookController.Create)
	router.GET("/api/books/:bookId", bookController.FindById)
	router.GET("/api/books", bookController.FindAll)
	router.PATCH("/api/books/:bookId", bookController.Update)
	router.DELETE("/api/books/:bookId", bookController.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}
