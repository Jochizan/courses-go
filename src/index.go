package main

import "fmt"

func main() {
	// Declaration const
	const pi float64 = 3.14
	const pi2 = 3.16

	fmt.Println("Pi:", pi)
	fmt.Println("Pi2:", pi2)

	// Declaration variables
	base := 12          // First version
	var height int = 14 // Second version
	var surface int     // Third version
	// no se compilara si las variables no son usadas

	fmt.Println(base, height, surface)

	// Zero values
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	// Area Cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado

	fmt.Println("El area del cuadrado es:", areaCuadrado)
}
