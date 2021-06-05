package main

import (
	"log"
	"net/http"
)

func main() {
	//Gorilla mux nos permite trabajar con un sistema de rutas implementando un router
	router := NewRouter()

	server := http.ListenAndServe(":8080", router)
	log.Fatal(server)
}
