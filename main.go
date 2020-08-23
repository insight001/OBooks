package main

import (
	"fmt"
	"log"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	log.Println("Route initiated")
	fmt.Println(r.Headers())
}
