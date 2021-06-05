package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var movies = Movies{
	Movie{Name: "Interstellar", Year: 2014, Director: "Christopher Nolan"},
	Movie{Name: "Avatar", Year: 2009, Director: "James Cameron"},
	Movie{Name: "Avengers: Endgame", Year: 2018, Director: "Russo's brothers"},
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola mundo desde mi servidor web con GO")
}

func MovieList(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(movies)
}

func MovieShow(w http.ResponseWriter, r *http.Request) {
	//De esta manera obtengo parametros de la url
	params := mux.Vars(r)
	movieId := params["id"]

	fmt.Fprintf(w, "Has cargado la pelicula número %s", movieId)
}

func MovieAdd(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var movieData Movie
	err := decoder.Decode(&movieData)

	if err != nil {
		panic(err)
	}

	defer r.Body.Close()

	log.Println(movieData)
	movies = append(movies, movieData)
}
