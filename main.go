package Movie_API

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

var movies []Movie

func getMovies(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(writer).Encode(movies)
	if err != nil {
		return
	}
}

func deleteMovie(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)

	for index, movie := range movies {
		if movie.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(writer).Encode(movies)
}

func getMovie(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	for _, movie := range movies {
		if movie.ID == params["id"] {
			json.NewEncoder(writer).Encode(movie)

		}
	}
}

func createMovie(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(request.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(100))
	movies = append(movies, movie)

	json.NewEncoder(writer).Encode(movie)
}

func updateMovie(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	for index, movie := range movies {
		if movie.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(request.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(writer).Encode(movie)
		}
	}

}

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "001", Title: "Dead Poets Society", Director: &Director{FirstName: "Peter", LastName: "Weir"}})
	movies = append(movies, Movie{ID: "2", Isbn: "002", Title: "Interstellar", Director: &Director{FirstName: "Christopher", LastName: "Nolan"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies{id}", deleteMovie).Methods("DELETE")

	fmt.Println("Movie API is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
