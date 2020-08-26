package data

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/insight001/OBooks/config"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq" //Databse config
)

//Database operations will be performed here

func goDotEnvVariable() string {

	// load .env file
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	conf := config.New()

	return conf.Database.URL
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

	fmt.Println(book)
	db, err := sql.Open("postgres", goDotEnvVariable())
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
func GetBooks(skip int, limit int, search string) ([]BookData, int) {

	db, err := sql.Open("postgres", goDotEnvVariable())
	if err != nil {
		log.Fatal("Failed to open a DB connection: ", err)
	}
	defer db.Close()

	if limit == 0 {
		limit = 10
	}

	fmt.Println("limit:", limit)
	fmt.Println("skip:", skip)
	fmt.Println("search:", search)
	query := `SELECT title,authors,isbn, description, id from books WHERE title LIKE '%' ||$1|| '%' offset $2 limit $3`

	rows, err := db.Query(query, search, skip, limit)

	if err != nil {
		// Do something
		panic(err)
	}
	defer rows.Close()

	var store []BookData

	for rows.Next() {
		var book BookData

		err = rows.Scan(&book.Title, &book.Authors, &book.ISBN, &book.Description, &book.ID)
		if err != nil {
			// handle this error
			panic(err)
		}

		store = append(store, book)
	}

	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		panic(err)
	}

	var count int

	err = db.QueryRow("SELECT COUNT(*) FROM books").Scan(&count)

	return store, count
}

//GetBook returns all books
func GetBook(id int) BookData {

	db, err := sql.Open("postgres", goDotEnvVariable())
	if err != nil {
		log.Fatal("Failed to open a DB connection: ", err)
	}
	defer db.Close()

	query := `Select id, title, authors, description, isbn from books where id=$1;`

	rows := db.QueryRow(query, id)
	if err != nil {
		// Do something
		panic(err)
	}

	var book BookData

	var bookID int
	var title, authours, isbn, description string
	err = rows.Scan(&bookID, &title, &authours, &description, &isbn)

	book.Description = description
	book.Authors = authours
	book.ID = bookID
	book.ISBN = isbn
	book.Title = title

	if err != nil {
		panic(err)
	}
	fmt.Println("getting books")

	return book
}
