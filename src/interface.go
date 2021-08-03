package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Employee struct {
	id int
}

type FullTimeEployee struct {
	Person
	Employee
	endDate string
}

// metodo de FullTimeEmployee
func (ftEmploye FullTimeEployee) getMessage() string {
	return "Full Time Employe"
}

type TemporaryEmployee struct {
	Person
	Employee
	taxRate int
}

// Metodo de Temporary Employee
func (tEmploye TemporaryEmployee) getMessage() string {
	return "Temporary Employe"
}

// Interface
type PrintInfo interface {
	getMessage() string
}

func getMessage(p PrintInfo) {
	fmt.Println(p.getMessage())
}

func main() {
	ftEmployee := FullTimeEployee{}
	ftEmployee.id = 1
	ftEmployee.name = "Maria"
	ftEmployee.age = 27
	//fmt.Printf("%v", ftEmployee)

	tEmployee := TemporaryEmployee{}
	getMessage(ftEmployee)
	getMessage(tEmployee)

}
