package loader

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/insight001/OBooks/config"
	"github.com/joho/godotenv"
)

//BookData ...
type BookData struct {
	ID          string   `json:"book_id"`
	Title       string   `json:"title"`
	Authors     []string `json:"authors"`
	ISBN        string   `json:"isbn"`
	Description string   `json:"description"`
}

type bookResponse struct {
	id          string
	authors     []string
	description string
	isbn        string
	title       string
}

//Get returns the book
func get() []BookData {
	fmt.Println("1. Performing Http Get...")
	resp, err := http.Get("https://learning.oreilly.com/api/v2/search/?fields=title&fields=isbn&fields=authors&limit=10")
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// Convert response body to an array of book struct
	var result map[string]interface{}

	//response := result["results"].(map[string][]interface{})

	err = json.Unmarshal(bodyBytes, &result)
	if err != nil {
		log.Fatalln(err)
	}

	s := make([]BookData, 10)
	for _, value := range result["results"].([]interface{}) {

		var book BookData
		book.Title = value.(map[string]interface{})["title"].(string)

		if _, ok := value.(map[string]interface{})["description"].(string); ok {
			book.Description = value.(map[string]interface{})["description"].(string)
		}

		r := value.(map[string]interface{})["authors"].([]interface{})

		for _, val := range r {
			book.Authors = append(book.Authors, val.(string))

		}
		book.ISBN = value.(map[string]interface{})["isbn"].(string)

		s = append(s, book)

	}

	return s
}
func goDotEnvVariable() string {

	// load .env file
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	conf := config.New()

	return conf.Database.URL
}

//BulkInsert loads
func BulkInsert() error {

	unsavedRows := get()

	db, err := sql.Open("postgres", goDotEnvVariable())
	if err != nil {
		log.Fatal("Failed to open a DB connection: ", err)
	}
	defer db.Close()

	valueStrings := make([]string, 0, len(unsavedRows))
	valueArgs := make([]interface{}, 0, len(unsavedRows)*4)
	i := 0
	for _, post := range unsavedRows {

		if post.Title != "" {
			fmt.Println(post)
			valueStrings = append(valueStrings, fmt.Sprintf("($%d, $%d, $%d,$%d)", i*4+1, i*4+2, i*4+3, i*4+4))
			valueArgs = append(valueArgs, post.Title)
			valueArgs = append(valueArgs, strings.Join(post.Authors, ";"))
			valueArgs = append(valueArgs, post.ISBN)
			valueArgs = append(valueArgs, post.Description)
			i++

		}

	}
	stmt := fmt.Sprintf("INSERT INTO books (Title, Authors, ISBN, Description) VALUES %s", strings.Join(valueStrings, ","))

	fmt.Println(stmt)
	fmt.Println(valueArgs...)
	_, err = db.Exec(stmt, valueArgs...)
	return err
}
