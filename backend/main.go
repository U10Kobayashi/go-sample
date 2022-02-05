package main

import (
    "log"
    "encoding/json"
    "fmt"
    "github.com/gorilla/mux"
    "net/http"
)

type Book struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Author *Author `json:"author"`
}

type Author struct {
    FirstName string `json:"firstname"`
    LastName  string `json:"lastname"`
}

func getBooks(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(books)
}

var books []Book

func main() {
    r := mux.NewRouter()

    books = append(books, Book{ID: "1", Title: "Book one", Author: &Author{FirstName: "Philip", LastName: "Williams"}})
    books = append(books, Book{ID: "2", Title: "Book Two", Author: &Author{FirstName: "John", LastName: "Johnson"}})

    r.HandleFunc("/api/books", getBooks).Methods("GET")
    // r.HandleFunc("/api/books/{id}", getBook).Methods("GET")

    fmt.Println("ok")

    log.Fatal(http.ListenAndServe(":8080", r))
}
