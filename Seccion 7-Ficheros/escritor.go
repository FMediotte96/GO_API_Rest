package main

import (
	"fmt"
	"os"
)

func main() {
	nuevo_texto := os.Args[1] + "\n"

	//err := ioutil.WriteFile("fichero.txt", nuevo_texto, 0777)

	fichero, err := os.OpenFile("fichero.txt", os.O_APPEND|os.O_WRONLY, 0777)
	showError2(err)

	escribir, err := fichero.WriteString(nuevo_texto)
	fmt.Println(escribir)
	showError2(err)

	err = fichero.Close()
	if err != nil {
		return
	}

}

func showError2(e error) {
	if e != nil {
		panic(e)
	}
}
