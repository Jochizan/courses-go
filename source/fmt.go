package main

import "fmt"

func main() {
	// Declaración de variables
	helloMessage := "Hello"
	worldMessage := "world"

	// Println
	fmt.Println(helloMessage, worldMessage)
	fmt.Println(helloMessage, worldMessage)

	// Printf
	name := "Platzi"
	courses := 500

	fmt.Printf("%s tiene más de %d cursos\n", name, courses)
	fmt.Printf("%v tiene más de %v cursos\n", name, courses)

	// Sprintf
	message := fmt.Sprintf("%s tiene más de %d cursos\n", name, courses)
	fmt.Println(message)

	// Tips
	fmt.Printf("HelloMessage: %T\n", helloMessage)
	fmt.Printf("cursos: %T\n", courses)
}
