package data

//BookStore contains method definition for the operations  performed on books
type BookStore interface {
	Initialize() //Load data from source
	SearchAuthor(author string, ratingOver, ratingBelow float64, limit, skip int) *[]*BookData
	SearchBook(bookName string, ratingOver, ratingBelow float64, limit, skip int) *[]*BookData
	SearchISBN(isbn string) *BookData
	CreateBook(book *BookData) bool
	DeleteBook(isbn string) bool
	UpdateBook(isbn string, book *BookData) bool
}
