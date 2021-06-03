package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now()
	day := now.Weekday()

	switch day {
	case 0:
		fmt.Println("Hoy es Domingo")
	case 1:
		fmt.Println("Hoy es Lunes")
	case 2:
		fmt.Println("Hoy es Martes")
	case 3:
		fmt.Println("Hoy es Miercoles")
	case 4:
		fmt.Println("Hoy es Jueves")
	default:
		fmt.Println("Es otro día de la semana")
	}
}
