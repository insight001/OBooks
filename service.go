package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/insight001/OBooks/data"
)

//SingleAPIResponse ...
type SingleAPIResponse struct {
	Success bool          `json:"success"`
	Message string        `json:"message"`
	Data    data.BookData `json:"data"`
}

//ListAPIResponse ...
type ListAPIResponse struct {
	Success bool            `json:"success"`
	Message string          `json:"message"`
	Data    []data.BookData `json:"data"`
}

//GetBookByID ..
func GetBookByID(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]
	val, err := strconv.Atoi(id)
	responseData := data.GetBook(val)
	var apiResponse SingleAPIResponse
	apiResponse.Data = responseData
	apiResponse.Success = true
	apiResponse.Message = "Items retrieved successfully"
	b, err := json.Marshal(apiResponse)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "error marshalling data"}`))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(b)
	return
}

//GetBooks ...
func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	limit, err := getLimitParam(r)
	skip, err := getSkipParam(r)
	search, err := getSearchParam(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "invalid datatype for parameter"}`))
		return
	}

	responseData := data.GetBooks(skip, limit, search)
	var apiResponse ListAPIResponse
	apiResponse.Data = responseData
	apiResponse.Success = true
	apiResponse.Message = "Items retrieved successfully"
	b, err := json.Marshal(apiResponse)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "error marshalling data"}`))
		return
	}
	w.WriteHeader(http.StatusOK)

	w.Write(b)
	return

}

func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "body not parsed"}`))
		return
	}

	var dataBook = data.BookData{Title: r.FormValue("Title"), Description: r.FormValue("Description"), Authors: r.FormValue("Authors"), ISBN: r.FormValue("ISBN")}

	ok := data.CreateBook(&dataBook)

	if ok {
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"success": "created"}`))
		return
	}
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"error": "not created"}`))
}

func getLimitParam(r *http.Request) (int, error) {
	limit := 0
	queryParams := r.URL.Query()
	l := queryParams.Get("limit")
	if l != "" {
		val, err := strconv.Atoi(l)
		if err != nil {
			return limit, err
		}
		limit = val
	}
	return limit, nil
}

func getSearchParam(r *http.Request) (string, error) {

	queryParams := r.URL.Query()
	l := queryParams.Get("limit")

	return l, nil
}

func getSkipParam(r *http.Request) (int, error) {
	skip := 0
	queryParams := r.URL.Query()
	l := queryParams.Get("skip")
	if l != "" {
		val, err := strconv.Atoi(l)
		if err != nil {
			return skip, err
		}
		skip = val
	}
	return skip, nil
}
