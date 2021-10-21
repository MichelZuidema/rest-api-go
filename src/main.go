package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type book struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
}

type allBooks []book

var books = allBooks{
	{
		Id:          1,
		Title:       "My First Book!",
		Author:      "John Doe",
		Description: "Testing 1 2 3",
	},
}

func createBook(rw http.ResponseWriter, r *http.Request) {
	var newBook book
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(rw, "No input")
	}

	json.Unmarshal(reqBody, &newBook)
	books = append(books, newBook)
	rw.WriteHeader(http.StatusCreated)

	json.NewEncoder(rw).Encode(newBook)
}

func updateBook(rw http.ResponseWriter, r *http.Request) {
	bid := mux.Vars(r)["id"]
	var newBook book

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(rw, "No input")
	}

	json.Unmarshal(reqBody, &newBook)

	for i, b := range books {
		bookId, err := strconv.Atoi(bid)

		if err != nil {
			fmt.Fprintf(rw, "Something went wrong")
		}

		if b.Id == bookId {
			b.Title = newBook.Title
			b.Description = newBook.Description
			b.Author = newBook.Author

			books = append(books[:i], b)

			json.NewEncoder(rw).Encode(b)
		}
	}
}

func deleteBook(rw http.ResponseWriter, r *http.Request) {
	bid := mux.Vars(r)["id"]

	for i, b := range books {
		bookId, err := strconv.Atoi(bid)

		if err != nil {
			fmt.Fprintf(rw, "Something went wrong")
		}

		if b.Id == bookId {
			books = append(books[:i], books[i+1:]...)
			fmt.Fprintf(rw, "The book with ID %v has been deleted successfully!", bookId)
		}
	}
}

func getBookById(rw http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	for _, el := range books {
		bookId, err := strconv.Atoi(id)

		if err == nil {
			if el.Id == bookId {
				json.NewEncoder(rw).Encode(el)
			}
		}
	}
}

func getAllBooks(rw http.ResponseWriter, r *http.Request) {
	json.NewEncoder(rw).Encode(books)
}

func welcome(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "Homescreen")
}

func main() {
	port := 8080

	router := mux.NewRouter().StrictSlash(true)
	fmt.Println("Listening on port 8080")

	router.HandleFunc("/", welcome).Methods("GET")

	router.HandleFunc("/books", getAllBooks).Methods("GET")
	router.HandleFunc("/book/{id}", getBookById).Methods("GET")
	router.HandleFunc("/book/{id}", updateBook).Methods("PATCH")
	router.HandleFunc("/book/{id}", deleteBook).Methods("DELETE")
	router.HandleFunc("/addBook", createBook).Methods("POST")

	log.Fatal(http.ListenAndServe(fmt.Sprint(":", port), router))
}
