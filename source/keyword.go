package main

import "fmt"

func main() {
	// Defer
	defer fmt.Println("Hola") // ejecutara esto antes que se acabe el programa
	fmt.Println("Mundo")

	for i := 0; i < 10; i++ {
		// continue
		if i == 2 {
			continue
		}

		// Break
		fmt.Println(i)

		if i == 8 {
			fmt.Println("Break")
			break
		}
	}
}
