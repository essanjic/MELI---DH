package main

import (
	"errors"
	"fmt"
)

/* Creamos un nuevo error personalizado utilizando la función errors.New() */
var err = errors.New("error: el salario es menor a 10.000")

/*
Ejercicio 3 - Impuestos de salario #3

Hacé lo mismo que en el ejercicio anterior pero reformulando el código para que, en reemplazo de “Error()”,  se implemente “errors.New()”.
*/
func main() {
	var salary int = 9000 // Creamos una variable salary con valor 9000
	// Comprobamos si el valor de salary es menor o igual a 10000
	if salary <= 10000 {
		// Usamos la función Is() de la librería errors para comprobar si el error es de tipo error
		if errors.Is(err, errors.New("")) {
			fmt.Println(err) // Imprimimos el mensaje de error en la consola
		}
	}
}
