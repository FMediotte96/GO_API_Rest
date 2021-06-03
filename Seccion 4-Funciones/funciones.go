package main

import "fmt"

func main() {
	/*holaMundo()

	var numero1 float32 = 10
	var numero2 float32 = 6

	fmt.Println("Calculadora 1")
	calculadora(numero1, numero2)

	var numero3 float32 = 44
	var numero4 float32 = 7

	fmt.Println("-------------")
	fmt.Println("Calculadora 2")
	calculadora(numero3, numero4)*/

	fmt.Println(devolverTexto2())
}

//palabra reservada func nombre_fn tipo_retorno
func holaMundo() {
	fmt.Println("Hola Mundo")
}

//puedo retornas más de un dato
func devolverTexto() (string, int) {
	dato1 := "Facundo"
	dato2 := 25
	return dato1, dato2
}

/*
	puedo definir el nombre de la variable en el retorno
	go reconoce que tiene devolver las variables que cree en el retorno de la fn
	de esta manera escribimos menos código y optimizamos el mismo
*/
func devolverTexto2() (dato1 string, dato2 int) {
	dato1 = "Facundo"
	dato2 = 25
	return
}

func operacion(n1 float32, n2 float32, op string) float32 {
	var resultado float32
	if op == "+" {
		resultado = n1 + n2
	}

	if op == "-" {
		resultado = n1 - n2
	}

	if op == "*" {
		resultado = n1 * n2
	}

	if op == "/" {
		resultado = n1 / n2
	}

	return resultado
}

func calculadora(numero1 float32, numero2 float32) {

	fmt.Print("La suma es: ")
	fmt.Println(operacion(numero1, numero2, "+"))

	fmt.Print("La resta es: ")
	fmt.Println(operacion(numero1, numero2, "-"))

	fmt.Print("La multiplicación es: ")
	fmt.Println(operacion(numero1, numero2, "*"))

	fmt.Print("La division es: ")
	fmt.Println(operacion(numero1, numero2, "/"))

}
