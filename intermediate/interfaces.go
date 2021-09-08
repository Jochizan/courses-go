package main

import "fmt"

type Person struct {
	age int
	name string
}

type Employee struct {
	id int
}

type FullTimeEmployee struct {
	Person
	Employee
}

type TemporaryEmployee struct {
	Person
	Employee
	taxRate int
}

type PrintInfo interface {
	getMessage() string
}

func (ftEmployee FullTimeEmployee) getMessage() string {
	return "FullTimeEmployee"
}

func (tEmployee TemporaryEmployee) getMessage() string {
	return "Temporary Employee"
}

func getMessage(p PrintInfo) {
	fmt.Printf(p.getMessage())
}

func main() {
	ftEmployee := FullTimeEmployee{}
	ftEmployee.name = "Jochizan"
	ftEmployee.age = 2
	ftEmployee.id = 5
	//fmt.Printf("%v\n", ftEmployee)

	tEmployee := TemporaryEmployee{}
	getMessage(tEmployee)
	getMessage(ftEmployee)
}
