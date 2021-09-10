package utils

import "fmt"

// CarPublic Car con acceso publico
type CarPublic struct {
	Brand string
	year  int
}

type carPrivate struct {
	brand string
	year  int
}

// PrintMessage imprimir un mensaje
func PrintMessage(text string) {
	fmt.Println(text)
}

func printMessageOne(text string) {
	fmt.Println(text)
}
