package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Movie struct {
	Id       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

var movies []Movie

func main() {
	r := mux.NewRouter()
	movies = append(movies, Movie{
		Id:       "1",
		Isbn:     "438227",
		Title:    "Movie 1",
		Director: &Director{FirstName: "John", LastName: "Doe"},
	})

	movies = append(movies, Movie{
		Id:       "2",
		Isbn:     "45455",
		Title:    "Movie 2",
		Director: &Director{FirstName: "Steve", LastName: "Smith"},	
	})
	r.HandleFunc("/movies").Methods("GET").HandlerFunc(getMovies)
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")


	fmt.Println("Starting server on port 8000")
	log.Fatal(http.ListenAndServe, r)

}