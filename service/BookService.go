package service

import (
	"github.com/mezz-ir/bookstore-clean-arch/interface" 
)

type BookService struct{
	BookDatalayer interface.BookDatalayer 

}

func NewBookService (BookDatalayer interface.BookDatalayer) interface.BookService{
	return &BookService{
		BookDatalayer : BookDatalayer
	}
}
 
func(service *BookService) PrintBookTitle (ctx context.Context, book *model.Book){

}
