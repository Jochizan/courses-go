package main

import "fmt"

func main() {

	// for conditional
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	fmt.Println()

	// For while
	counter := 0
	for counter <= 10 {
		fmt.Println(counter)
		counter++
	}

	// For forever
	// counterForever := 0
	// for {
	// 	fmt.Println(counterForever)
	// 	counterForever++
	// }

	listOdds := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
	for i, par := range listOdds {
		fmt.Printf("posicion %d número par: %d \n", i, par)
	}

	for i := 10; i > 0; i-- {
		fmt.Println(i)
	}
}
