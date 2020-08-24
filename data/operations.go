package data

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

//Database operations will be performed here

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

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

	db, err := sql.Open("petextmt", goDotEnvVariable("DB_URL"))
	if err != nil {
		log.Fatal("Failed to open a DB connection: ", err)
	}
	defer db.Close()

	query := "INSERT INTO books (Title, Authors, ISBN, Description) VALUES ($1, $2, $3,$4) RETURNING id"
	id := 0
	err = db.QueryRow(query, book.Title, book.Authors, book.ISBN, book.Description).Scan(&id)
	if err != nil {
		panic(err)
	}
	fmt.Println("New record ID is:", id)
	fmt.Println("Book Created")
	return true
}

//GetBooks returns all books
func GetBooks(skip, limit int, search string) []BookData {

	db, err := sql.Open("petextmt", goDotEnvVariable("DB_URL"))
	if err != nil {
		log.Fatal("Failed to open a DB connection: ", err)
	}
	defer db.Close()

	if limit == 0 {
		limit = 10
	}
	query := "Select * from books offset $0"

	if search != "" {
		query = "Select * from books offset $0 limit $1 where description like %$2%"
	} else {
		query = "Select * from books offset $0 limit $1"
	}

	rows, err := db.Query(query, skip, limit, search)

	if err != nil {
		// Do something
		panic(err)
	}
	defer rows.Close()

	store := make([]BookData, limit)

	for rows.Next() {
		var book BookData

		err = rows.Scan(&book.ID, &book.Title, &book.Authors, &book.ISBN, &book.Description)
		if err != nil {
			// handle this error
			panic(err)
		}
		fmt.Println(&book.Description)
		store = append(store, book)
	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	fmt.Println("getting books")
	return store
}
