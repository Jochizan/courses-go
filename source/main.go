package main

import (
	pk "platzi.est/m/mypackage"
	"fmt"
)

func main() {
	var myCar pk.CarPublic
	myCar.Brand = "Ferrari"
	//myCar.Year = 2020
	fmt.Println(myCar)

	//var myOtherCar pk.carPrivate
	//fmt.Println(myOtherCar)
	pk.PrintMessage("Hola Platzi")
	pk.printMessageOne("Hola")
}
