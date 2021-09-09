package main

import "fmt"

func sum(values ...int) int {
	total := 0
	for _, num := range values {
		total += num
	}
	return total
}

// recibe un slice de valores string y los imprime.
func printNames(names ...string) {
	for _, name := range names {
		fmt.Println(name)
	}
}

// Funcion que retorna multiples valores.
func getValues(x int) (double int, triple int, quad int) {
	// return 2 * x, 3 * x, 4 * x // vieja forma.
	double = 2 * x
	triple = 3 * x
	quad = 4 * x
	return
}

func main() {
	fmt.Println(sum(1, 2, 3))
	printNames("joan", "jose", "armando", "luis")

	// Imprimiento multiples returns:
	fmt.Println(getValues(5))

	// capturando multiples returns:
	d, t, q := getValues(5)
	fmt.Println(d, t, q)
}

