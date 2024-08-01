package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

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


func getMovies (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, movie := range movies {
		if movie.Id == params["id"] {
			movies = append(movies[:index],  movies[index + 1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func createMovie (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.Id = strconv.Itoa(len(movies) + 1)
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movies)
}

func getMovie (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, movie := range movies {
		if movie.Id == params["id"] {
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}


func updateMovie (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, movie := range movies {
		if movie.Id == params["id"] {
			var movie Movie
			movies = append(movies[:index], movies[index + 1:]...)
			json.NewDecoder(r.Body).Decode(&movie)
			movie.Id = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movies)
		}
	}
}


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
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")


	fmt.Println("Starting server on port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))

}