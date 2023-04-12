package main

import "fmt"

func main() {
	// 	Ejercicio 1 - Letras de una palabra
	// La Real Academia Española quiere saber cuántas letras tiene una palabra y luego tener cada una de las letras por separado para deletrearla. Para eso tendrán que:
	// Crear una aplicación que reciba  una variable con la palabra e imprimir la cantidad de letras que contiene la misma.
	// Luego, imprimí cada una de las letras.

	palabra := []string{"holo"}
	cantidadLetras := len(palabra[0])
	fmt.Println("La cantidad de letras que tiene la palabra son:", cantidadLetras)

	//imprimir cada una de las letras
	for i := 0; i < cantidadLetras; i++ {
		fmt.Println("Las letras de la palabra son:", string(palabra[0][i]))
	}
}
