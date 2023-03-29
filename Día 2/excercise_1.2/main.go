package main

import (
	"fmt"
)

func main() {
	var palabra = []string{"holas"}
	cantidadLetras := len(palabra[0])

	fmt.Println("La cantidad de letras que tiene la palabra son:", cantidadLetras)
	fmt.Println(palabra[0][0:])
}
