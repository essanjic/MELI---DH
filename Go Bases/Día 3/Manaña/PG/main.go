package main

import (
	"fmt"
)

func main() {
	var a, b int
	fmt.Println("Ingrese los dos números a sumar:")
	fmt.Scan(&a, &b)
	resultado := suma(a, b)
	fmt.Println("El resultado es: ", resultado)

}

func suma(a, b int) int {
	return a + b

}
