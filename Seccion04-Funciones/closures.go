package main

import "fmt"

func main() {
	fmt.Print("Pedido 1 -->")
	fmt.Println(gorras(45, "€"))

	fmt.Println("--------------")

	fmt.Print("Pedido 2 -->")
	fmt.Println(gorras(24, "usd"))
}

/*
	Dentro de go podemos crear una fn y dentro de la misma crear otra fn, esto es conocido
	como closure, lo que nos permite definir dentro de una fn una variable en la cual dentro
	hay una fn anonima
*/
func gorras(pedido float32, moneda string) (string, float32, string) {
	//declaro una fn anonima
	precio := func() float32 {
		return pedido * 7
	}

	return "El precio del pedido es:", precio(), moneda
}
