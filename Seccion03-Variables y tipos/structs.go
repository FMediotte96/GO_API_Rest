package main

import "fmt"

type Gorra struct {
	marca  string
	color  string
	precio float32
	plana  bool
}

func main() {

	//se puede indicar el nombre del atributo o inicializarlo con el orden declarado
	var gorra_negra = Gorra{"Adidas", "Roja", 20.50, true}

	//Accedo a sus atributos con ".<nombre_atributo>"
	fmt.Println(gorra_negra.marca)
	fmt.Println(gorra_negra.precio)
}
