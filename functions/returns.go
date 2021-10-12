package main

import "fmt"

func suma(valores ...int) int {
	total := 0
	for _, numero := range valores {
		total += numero
	}
	return total
}

func getValores(x int) (doble int, triple int, cuatro int) {
	doble = x * 2
	triple = x * 3
	cuatro = x * 4
	return
}

func main() {
	fmt.Println(suma(1, 2))

	fmt.Println(suma(1, 2, 3, 4))

	fmt.Println(getValores(3))
}
