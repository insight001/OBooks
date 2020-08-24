package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/insight001/OBooks/data"
)

var (
	books data.BookData
)

func main() {
	r := mux.NewRouter()
	log.Println("bookdata api")
	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "api v1")
		log.Println("Got here")
	})
	api.HandleFunc("/books", GetBooks).Methods(http.MethodGet)
	api.HandleFunc("/book", createBook).Methods(http.MethodPost)
	log.Fatalln(http.ListenAndServe(":8080", r))
}
