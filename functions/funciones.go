package main

import (
	"fmt"
	"time"
)

func main() {
	x := 5
	y := func() int {
		return x * 2
	}()

	fmt.Printf("[i] El doble de %d es %d\n\n", x, y)
	canal := make(chan int)
	func() {
		fmt.Println("[+] Funcion larga")
		time.Sleep(5 * time.Second)
		canal <- 1
		fmt.Println("[i] Termine funcion larga")
	}()
	<-canal
}
