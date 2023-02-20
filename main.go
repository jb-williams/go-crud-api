package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `jsong:"director"`
}
type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func getMovies(resp http.ResponseWriter, requ *http.Request) {
	resp.Header().Set("Content-Type", "application/json")
	json.NewEncoder(resp).Encode(movies)
}

func getMovie(resp http.ResponseWriter, requ *http.Request) {
	resp.Header().Set("Content-Type", "application/json")
	params := mux.Vars(requ)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(resp).Encode(item)
			return
		}
	}
}

func createMovie(resp http.ResponseWriter, requ *http.Request) {
	resp.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(requ.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(100000000))
	movies = append(movies, movie)
	json.NewEncoder(resp).Encode(movie)
}

// not best for working with databases
func updateMovie(resp http.ResponseWriter, requ *http.Request) {
	// set json content type
	resp.Header().Set("Content-Type", "application/jsson")
	// params
	params := mux.Vars(requ)
	// loop over the movies, range
	// delete the movie with the id that was sent
	// add new movie - the movie that was sent in body
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(requ.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(resp).Encode(movie)
			return
		}
	}

}

func deleteMovie(resp http.ResponseWriter, requ *http.Request) {
	resp.Header().Set("Content-Type", "application/json")
	params := mux.Vars(requ)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(resp).Encode(movies)
}

func main() {
	// mux := http.NewServeMux()
	router := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "438227", Title: "Movie One", Director: &Director{Firstname: "John", Lastname: "Doe"}})
	movies = append(movies, Movie{ID: "2", Isbn: "45455", Title: "Movie Two", Director: &Director{Firstname: "Steve", Lastname: "Smith"}})
	router.HandleFunc("/movies", getMovies).Methods("GET")
	router.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	router.HandleFunc("/movies", createMovie).Methods("POST")
	router.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	router.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")
	fmt.Printf("Starting Server at localhost:9999\n")
	err := http.ListenAndServe(":9999", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		// fmt.Printf("error listening for server: %s\n", err)
		log.Fatal(err)
	}
}
