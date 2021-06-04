package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	//Gorilla mux nos permite trabajar con un sistema de rutas implementando un router
	router := mux.NewRouter().StrictSlash(true) //indica que las urls son amigables utilizando "/"
	router.HandleFunc("/", Index)
	router.HandleFunc("/contacto", Contact)

	server := http.ListenAndServe(":8080", router)
	log.Fatal(server)
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola mundo desde mi servidor web con GO")
}

func Contact(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Esta es la página de contacto")
}
