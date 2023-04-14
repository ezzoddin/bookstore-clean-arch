package interface

import (
	"github.com/mezz-ir/bookstore-clean-arch/model" 
	"context"
)

type BookService interface {
	PrintBookTitle(ctx context.Context, book *model.Book)
}

type BookDatalayer interface {
	GetBookByID(ctx context.Context, bookID int16)
}