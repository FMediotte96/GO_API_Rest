package main

import "fmt"

func main() {
	var peliculas [3][2]string
	peliculas[0][0] = "Inception"
	peliculas[0][1] = "Interstellar"
	peliculas[1][0] = "Fight Club"
	peliculas[1][1] = "The Lord of the Rings"
	peliculas[2][0] = "Ironman"
	peliculas[2][1] = "Captain America"

	fmt.Println(peliculas[2][1])
}
