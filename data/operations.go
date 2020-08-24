package data

import "fmt"

//Database operations will be performed here

//BookStore contains method definition for the operations  performed on books
type BookStore interface {
	Initialize() //Load data from source
	GetBooks(search string, limit, skip int) *[]*BookData
	CreateBook(book *BookData) bool
}

//Books ...
type Books struct {
	Store *[]*BookData `json:"store"`
}

//CreateBook ...
func CreateBook(book *BookData) bool {

	fmt.Println("Book Created")
	return true
}

//GetBooks returns all books
func GetBooks(skip, limit int) *[]*BookData {

	fmt.Println("getting books")
	return nil
}
