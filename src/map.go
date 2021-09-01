package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Joan"] = 14
	m["Pepe"] = 20

	fmt.Println(m["Joser"])

	// Recorrer map
	for i, v := range m {
		fmt.Println(i, v)
	}

	// Encontrar valor
	value, ok := m["Joa"]
	fmt.Println(value, ok)
}
