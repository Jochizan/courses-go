package main

import (
	"fmt"
)

type figure2D interface {
	surface() float64
}

type square struct {
	base float64
}

type rectangle struct {
	width  float64
	height float64
}

func (c square) surface() float64 {
	return c.base * c.base
}

func (r rectangle) surface() float64 {
	return r.width * r.height
}

func calcule(f figure2D) {
	fmt.Println("Area: ", f.surface())
}

func main() {
	mySquare := square{ base: 2 }
	myRectangle := rectangle{ width: 2, height: 4 }

	calcule(mySquare)
	calcule(myRectangle)

	// Lista interfaces
	myInterface := []interface{}{"Hola", 12, 4.98}
	fmt.Println(myInterface...)
}
