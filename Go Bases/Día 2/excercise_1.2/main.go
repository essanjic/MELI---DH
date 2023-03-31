package main

import (
	"fmt"
)

func main() {
	var palabra = []string{"holas"}
	cantidadLetras := len(palabra[0])

	fmt.Println("La cantidad de letras que tiene la palabra son:", cantidadLetras)
	for i := 0; i < cantidadLetras; i++ {
		fmt.Println("La letra", i+1, "es:", palabra[0][i])
	}
}
