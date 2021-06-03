package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fmt.Println("Lector")

	texto, errFichero := ioutil.ReadFile("fichero.txt")
	showError(errFichero)

	fmt.Println(string(texto))
}

func showError(e error) {
	if e != nil {
		panic(e)
	}
}
