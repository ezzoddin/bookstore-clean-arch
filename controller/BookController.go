package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/mezz-ir/bookstore-clean-arch/interface" 
)

type BookController struct {
	bookService interface.BookService
}

func NewBookController(echoCtx $echo.Echo, bookService interface.BookService)  {
	bookControllerHandler := &BookController {
		bookService: bookService
	}

	echoCtx.Get("/printAuthor", bookControllerHandler.bookService)
}

func (bookController *BookController) printAuthor(ec echo.Context) error{
	return nil
}