package main

import "fmt"

func main() {
	var suma int = 8 + 9
	var resta int = 6 - 4
	var nombre string = "Facundo "

	var apellidos string = "Mediotte Palermo"
	apellidos = "Martinez "

	/*
		Puedo crear variables utilizando ":=" de esta forma go va a inferir
		el tipo de dato que va a tener la variable
	*/
	pais := "Argentina"

	var prueba bool = true
	fmt.Println(prueba)

	var f32 float32 = 12.5
	fmt.Println(f32)

	//Constante
	const year int = 2021
	fmt.Println(year)

	fmt.Println(suma)
	fmt.Println(resta)
	fmt.Println("Mi nombre es " + nombre + apellidos + pais)
}

/*
	Go es un lenguaje tipado por lo que debemos definir siempre un tipo de dato asignado a las variables
*/
