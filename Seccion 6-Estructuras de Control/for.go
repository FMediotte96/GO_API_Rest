package main

import "fmt"

func main() {

	tope := 26
	for i := 1; i <= tope; i++ {
		if i%2 == 0 {
			fmt.Println("El número es par:", i)
		} else {
			fmt.Println("El número es impar:", i)
		}
	}

	peliculas := []string{"Fight Club", "Batman", "Ironman"}
	//Recorriendo un slice, del estilo forEach
	for _, pelicula := range peliculas {
		fmt.Println(pelicula)
	}
}
