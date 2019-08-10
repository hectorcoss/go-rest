package main

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
)

type Book struct {
	ID int `json:id`
	Title string `json:title`
	Author string `json:author`
	Year string `json:year`
}

var books []Book

func main() {

	books = append(books, Book{ID: 1, Title: "Guerras de Dios", Author: "Christopher Tyreman", Year: "2018"})

	router := mux.NewRouter()
	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/books", addBook).Methods("POST")
	router.HandleFunc("/books", updateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", removeBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func getBooks(w http.ResponseWriter, r *http.Request) {
    log.Println("Gets all books")
    json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
    log.Println("Gets one book")
    params := mux.Vars(r)
    i, _ := strconv.Atoi(params["id"])
    
    for _, book := range books {
    	if book.ID == i {
    		json.NewEncoder(w).Encode(&book)
    	}
    }
}

func addBook(w http.ResponseWriter, r *http.Request) {
    log.Println("Adds a book")
	var book Book
	json.NewDecoder(r.Body).Decode(&book)

	books = append(books, book)

	json.NewEncoder(w).Encode(books)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
    log.Println("Updates a book")
    var book Book
	json.NewDecoder(r.Body).Decode(&book)

	for i, item := range books {
		if item.ID == book.ID {
			books[i] = book
		}
	}

	json.NewEncoder(w).Encode(books)
}

func removeBook(w http.ResponseWriter, r *http.Request) {
    log.Println("Deletes a book")
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])

	for i, item := range books {
    	if item.ID == id {
    		books = append(books[:i], books[i+1:]...)
    	}
    }

	json.NewEncoder(w).Encode(books)
}