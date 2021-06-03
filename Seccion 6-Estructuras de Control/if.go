package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("**** Mi programa con Go ****")

	//Argumento de consola
	fmt.Println("Hola " + os.Args[1] + " bienvenido al programa en go")
	edad, _ := strconv.Atoi(os.Args[2])

	if edad >= 18 && edad != 25 && edad != 99 {
		fmt.Println("Eres mayor de edad")
	} else if edad == 25 {
		fmt.Println("Tienes 25 años")
	} else if edad == 99 {
		fmt.Println("Eres muy viejo")
	} else {
		fmt.Println("Eres menor de edad")
	}
}

/*
	== igual a
	!= diferente de
	< menor que
	> mayor que
	>= mayor o igual que
	<= menor o igual que
	&& AND
	|| OR
*/
