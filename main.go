package main

import (
	"fmt"
	"net/http"
	"log"
	"github.com/gorilla/mux"
	"encoding/json"
	"math/rand"
	"strconv"
)


type Movie struct {
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}


type Director struct {
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
}


var movies []Movie {
	{ID: "123", Isbn: "456321", Title: "Harry Potter", Director: &Director{FirstName: "JK", LastName: "Rowling"}},
	{ID: "561", Isbn: "321654", Title: "HIT", Director: &Director{FirstName: "Rishika", LastName: "Gupta"}},
	{ID: "852", Isbn: "852649", Title: "Anabelle", Director: &Director{FirstName: "Bhumika", LastName: "Chaurasia"}}
}


func getMovies(w http.ResponseWritter, r *http.Request){

}


func getMovie(w http.ResponseWritter, r *http.Request){
	w.Header.Set("Content-type", "application/json")
	params := mux.Vars(r)
	for _, item in range movies {
		if item.Id == params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}


func createMovie(w http.ResponseWritter, r *http.Request){
	w.Header.Set("Content-type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(1000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}


func updateMovie(w http.ResponseWritter, r *http.Request){
	w.Header.Set("Content-type", "application/json")
	params := mux.Vars(r)

	for index, item in range movies {
		if item.Id == params["id"] {
			movies = append()
		}
	}
}


func deleteMovie(w http.ResponseWritter, r *http.Request){

}


func main(){
	r := mux.newRouter()
	
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting Server at 8000\n")
	log.Fatal(http.ListenAndServer(":8000", r))

}