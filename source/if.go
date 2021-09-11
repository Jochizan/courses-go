package main

import "fmt"

func main() {
	value1 := 1
	value2 := 2

	if value1 == 1 {
		fmt.Println("Es 1")
	} else {
		fmt.Println("No es 1")
	}

	// with and
	if value1 == 1 && value2 == 3 {
		fmt.Println("Es verdad, AND")
	}

	// with or
	if value1 == 0 || value2 == 2 {
		fmt.Println("Es verdad, OR")
	}
}
