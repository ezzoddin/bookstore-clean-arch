package database

type BookDatalayer struct{
}

func NewBookDatalayer () interface.BookDatalayer{
	return &BookDatalayer{}
}


func (bookDatalayer *BookDatalayer) GetBookByID(ctx context.Context, bookID int16) {

}
