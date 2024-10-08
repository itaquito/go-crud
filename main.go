package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

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

func getMovies(w http.ResponseWriter, r *http.Request) {
	// Define that we are going to return JSON
	w.Header().Set("Content-Type", "application/json")

	// Encode the movies as JSON and reply it
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	// Define that we are going to return JSON
	w.Header().Set("Content-Type", "application/json")

	// Extract the params of the request
	params := mux.Vars(r)

	// Find the movie with ID
	for index, item := range movies {
		if item.ID == params["id"] {
			// Delete a movie with append
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}

	// Returns the list of movies
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	// Define that we are going to return JSON
	w.Header().Set("Content-Type", "application/json")

	// Extract the params of the request
	params := mux.Vars(r)

	// Find the movie with ID
	// In this case, we use _ since we won't be using the index
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	// Define that we are going to return JSON
	w.Header().Set("Content-Type", "application/json")

	// A variable of type movie
	var movie Movie
	// Decode the body and save it to the movie variable
	_ = json.NewDecoder(r.Body).Decode(&movie)

	// Create a random ID and covert it into a string
	movie.ID = strconv.Itoa(rand.Intn(100))

	// Appends the new movie into our list
	movies = append(movies, movie)

	// Returns the list of movies
	json.NewEncoder(w).Encode(movies)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	// Define that we are going to return JSON
	w.Header().Set("Content-Type", "application/json")

	// Extract the params of the request
	params := mux.Vars(r)

	// Find the movie with ID
	for index, item := range movies {
		if item.ID == params["id"] {
			// Delete a movie with append
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}

	// A variable of type movie
	var movie Movie
	// Decode the body and save it to the movie variable
	_ = json.NewDecoder(r.Body).Decode(&movie)

	// Create a random ID and covert it into a string
	movie.ID = params["id"]

	// Appends the new movie into our list
	movies = append(movies, movie)

	// Returns the list of movies
	json.NewEncoder(w).Encode(movies)
}

func main() {
	// Creates the MUX router
	r := mux.NewRouter()

	// Create some default movies
	movies = append(movies, Movie{ID: "1", Isbn: "12345", Title: "Movie 1", Director: &Director{Firstname: "John", Lastname: "Doe"}})
	movies = append(movies, Movie{ID: "2", Isbn: "67890", Title: "Movie 2", Director: &Director{Firstname: "Peter", Lastname: "Smith"}})

	// Define the routes
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port 8000.\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}
