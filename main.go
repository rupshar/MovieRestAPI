package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Movie struct {
	Title string `json:"title"`
	Director *Person `json:"director"`
	Producer *Person `json:"producer"`
	Year int `json:"year"`
	MainActor *Person `json:"mainactor"`
	MainActress *Person `json:"mainactress"`
	SupportingActor *Person `json:"suppactor"`
	SupportingActress *Person `json:"suppactress"`
}

type Person struct {
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range movies {
		if item.Title == params["title"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Movie{})
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.Title == params["title"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.Title = params["title"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.Title == params["title"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		} 
	}
	json.NewEncoder(w).Encode(movies)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{title}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{title}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{title}", deleteMovie).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":5000", r))
}