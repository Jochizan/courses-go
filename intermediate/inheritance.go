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

func GetMessage(p Person) {
	fmt.Printf("%s with age %d\n", p.name, p.age)
}

func (employee FullTimeEmployee) String() string {
	return fmt.Sprintf("\nid: %d, name: %s, age: %d ", employee.id, employee.name, employee.age)
}

func main() {
	ftEmployee := FullTimeEmployee{}
	ftEmployee.name = "Name"
	ftEmployee.age = 2
	ftEmployee.id = 5

	fmt.Printf("%v\n", ftEmployee)
}
