package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// A struct that defines a Movie
type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

// A struct that defines a Director
type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// An array of Movies
var movies []Movie

func main() {
	// Creates the MUX router
	r := mux.NewRouter()

	// Create some default movies
	movies = append(movies, Movie{ID: "1", Isbn: "12345", Title: "Movie 1", Director: &Director{Firstname: "John", Lastname: "Doe"}})
	movies = append(movies, Movie{ID: "2", Isbn: "67890", Title: "Movie 2", Director: &Director{Firstname: "Peter", Lastname: "Smith"}})

	fmt.Printf("Starting server at port 8000.\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}
