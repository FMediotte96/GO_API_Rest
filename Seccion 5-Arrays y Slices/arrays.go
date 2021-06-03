package main

import "fmt"

func main() {
	/*
		var peliculas [3]string
		peliculas[0] = "Inception"
		peliculas[1] = "Interstellar"
		peliculas[2] = "Fight Club"
	*/

	//podemos definir un array de la siguiente manera tambien
	peliculas := [3]string{"Inception",
		"Interstellar",
		"Fight Club"}

	fmt.Println(peliculas[2])

}
