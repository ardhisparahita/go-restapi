package controller

import (
	"go-restapi/helper"
	"go-restapi/model/web"
	"go-restapi/service"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type BookControllerImpl struct {
	BookService service.BookService
}

func NewBookController(bookService service.BookService) BookController {
	return &BookControllerImpl{
		BookService: bookService,
	}
}

func (controller *BookControllerImpl) Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	bookCreateRequest := web.BookCreateRequest{}
	helper.ReadFromRequestBody(r, &bookCreateRequest)

	bookResponse := controller.BookService.Create(r.Context(), bookCreateRequest)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   bookResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *BookControllerImpl) FindById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	bookId := p.ByName("bookId")

	bookResponse := controller.BookService.FindById(r.Context(), bookId)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   bookResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *BookControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	bookResponses := controller.BookService.FindAll(r.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   bookResponses,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *BookControllerImpl) Update(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	bookUpdateRequest := web.BookUpdateRequest{}

	bookId := p.ByName("bookId")
	helper.ReadFromRequestBody(r, &bookUpdateRequest)

	bookUpdateRequest.Id = bookId

	bookResponse := controller.BookService.Update(r.Context(), bookUpdateRequest)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   bookResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *BookControllerImpl) Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	bookId := p.ByName("bookId")

	controller.BookService.Delete(r.Context(), bookId)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(w, webResponse)
}
