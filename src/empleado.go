package main

import "fmt"

type Employee struct {
	id       int
	name     string
	vacation bool
}

// Creacion de un constructor
func NewEmployee(id int, name string, vacation bool) *Employee {
	return &Employee{
		id:       id,
		name:     name,
		vacation: vacation,
	}
}

func (e *Employee) SetId(id int) {
	e.id = id
}

func (e *Employee) SetName(name string) {
	e.name = name
}

func (e *Employee) GetId() int {
	return e.id
}

func (e *Employee) GetName() string {
	return e.name
}

func main() {
	e := Employee{}
	fmt.Printf("Empleado 1: %v\n", e)
	e.id = 1
	e.name = "Ana"
	e.SetId(5)
	e.SetName("Luis")
	fmt.Println(e.GetId())
	fmt.Println(e.GetName())

	// forma numero 2.
	e2 := Employee{
		id:   1,
		name: "Jose",
	}
	fmt.Printf("Empleado 2: %v\n", e2)

	// forma 3.
	e3 := new(Employee)                 // new devuelve un apuntador
	fmt.Printf("Empleado 3: %v\n", *e3) // con * referenciamos al valor.

	// forma 4. Crear una funcion llamada New algo y retornar un apuntador
	e4 := NewEmployee(4, "Laura", false)
	fmt.Printf("Empleado 4: %v\n", *e4)
}
