// Package main provides ...
package main

import "fmt"

type pc struct {
	ram   int
	disk  int
	brand string
}

func (myPC pc) ping() {
	fmt.Println(myPC.brand, "Pong")
}

func (myPC *pc) duplicateRAM() {
	myPC.ram = myPC.ram * 2
}

func (myPC pc) String() string {
	return fmt.Sprintf("Tengo %d GB RAM, %d GB Disco y es una %s", myPC.ram, myPC.disk, myPC.brand)
}

func main() {
	
	a := 50
	b := &a

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(*b)

	//	cambiando el valor de la dirección de memoria
	*b = 100
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(*b)

	myPC := pc{ram: 16, disk: 200, brand: "msi"}

	fmt.Println(myPC)
	myPC.ping()
	
	myPC.duplicateRAM()
	fmt.Println(myPC)

	myPC.duplicateRAM()
	fmt.Println(myPC)
}
