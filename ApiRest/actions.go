package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
)

func getSession() *mgo.Session {
	session, err := mgo.Dial("mongodb://localhost")

	if err != nil {
		panic(err)
	}

	return session
}

func responseMovie(w http.ResponseWriter, status int, result Movie) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(result)
}

func responseMovies(w http.ResponseWriter, status int, results []Movie) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(results)
}

var collection = getSession().DB("curso_go").C("movies")

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola mundo desde mi servidor web con GO")
}

func MovieList(w http.ResponseWriter, r *http.Request) {
	var results []Movie
	err := collection.Find(nil).Sort("-_id").All(&results) //ordeno al revés por id

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Resultados: ", results)
	}

	responseMovies(w, 200, results)
}

func MovieShow(w http.ResponseWriter, r *http.Request) {
	//De esta manera obtengo parametros de la url
	params := mux.Vars(r)
	movieId := params["id"]

	if !bson.IsObjectIdHex(movieId) {
		w.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(movieId)

	result := Movie{}
	err := collection.FindId(oid).One(&result)

	if err != nil {
		w.WriteHeader(404)
		return
	}

	responseMovie(w, 200, result)
}

func MovieAdd(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var movieData Movie
	err := decoder.Decode(&movieData)

	if err != nil {
		panic(err)
	}

	defer r.Body.Close()

	err = collection.Insert(movieData)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	responseMovie(w, 200, movieData)
}

func MovieUpdate(w http.ResponseWriter, r *http.Request) {
	//De esta manera obtengo parametros de la url
	params := mux.Vars(r)
	movieId := params["id"]

	if !bson.IsObjectIdHex(movieId) {
		w.WriteHeader(404)
		return
	}
	oid := bson.ObjectIdHex(movieId)
	decoder := json.NewDecoder(r.Body)

	var movieData Movie
	err := decoder.Decode(&movieData)
	if err != nil {
		w.WriteHeader(500)
		panic(err)
		return
	}
	defer r.Body.Close()

	document := bson.M{"_id": oid}
	change := bson.M{"$set": movieData}
	err = collection.Update(document, change)

	if err != nil {
		w.WriteHeader(404)
		panic(err)
		return
	}

	responseMovie(w, 200, movieData)
}

type Message struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func (m *Message) setStatus(data string) {
	m.Status = data
}

func (m *Message) setMessage(message string) {
	m.Message = message
}

func MovieRemove(w http.ResponseWriter, r *http.Request) {
	//De esta manera obtengo parametros de la url
	params := mux.Vars(r)
	movieId := params["id"]

	if !bson.IsObjectIdHex(movieId) {
		w.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(movieId)

	err := collection.RemoveId(oid)

	if err != nil {
		w.WriteHeader(404)
		return
	}

	//result := Message{"success", "La pelicula con ID " + movieId + " ha sido eliminada correctamente"}

	message := new(Message)
	message.setStatus("success")
	message.setMessage("La pelicula con ID " + movieId + " ha sido eliminada correctamente")

	result := message
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(result)
}
