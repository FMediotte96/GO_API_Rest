package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	numero, _ := strconv.Atoi(os.Args[1])

	if numero%2 == 0 {
		fmt.Println("El número es par")
	} else {
		fmt.Println("El número es impar")
	}
}
