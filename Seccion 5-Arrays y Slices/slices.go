package main

import "fmt"

func main() {
	//Slices: estructura que tiene como base los arrays, pero en este caso tenemos una dimensión dinamica
	peliculas := []string{"Inception",
		"Interstellar",
		"Fight Club",
		"Superman"}

	peliculas = append(peliculas, "Batman")
	peliculas = append(peliculas, "Ironman")
	peliculas = append(peliculas, "Inception")

	//Funciones útiles
	fmt.Println(len(peliculas)) // me muestra la cantidad de elementos

	//Slicing (desde:hasta) no inclusivos
	fmt.Println(peliculas[0:3])
}
