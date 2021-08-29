package main

import "fmt"
import "strings"

func isPalindrome(text string) (bool) {
	var textReverse string
	text = strings.ToLower(text);
	
	for i := len(text) - 1; i >= 0; i-- {
		textReverse += string(text[i])
	}

	return text == textReverse
}

func main() {
	slice := []string{"hola", "que", "hace"}

	//for _, value := range(slice) {
		//fmt.Println(value)
	//}

	for value := range slice {
		fmt.Println(value)
	}

	fmt.Println(isPalindrome("Ama"))
}
