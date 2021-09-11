package main

import (
	"fmt"
	"math"
)

func areaCirculo(radio float64) float64 {
	return math.Pi * radio * radio
}

func areaRectangulo(base float64, altura float64) float64 {
	return base * altura
}

func areaTrapezoide(B float64, b float64, h float64) float64 {
	return h * (B + b) / 2
}

func doubleReturn(a int) (c, d int) {
	return a, a * 2
}

func main() {

	value1, value2 := doubleReturn(2)

	fmt.Println(value1, value2)
	fmt.Printf("Circulo %.2f \n", areaCirculo(2))
	fmt.Printf("Rectangulo %.2f \n", areaRectangulo(5, 10))
	fmt.Printf("Trapezoide %.2f \n", areaTrapezoide(10, 5, 3))

}
