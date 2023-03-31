package main

import (
	"errors"
	"fmt"
)

/* Ejercicio 2 - Impuestos de salario #2
En tu función “main”, definí una variable llamada “salary” y asignale un valor de tipo “int”.
Creá un error personalizado con un struct que implemente “Error()” con el mensaje “Error: el salario es menor a 10.000" y lanzalo en caso de que “salary” sea menor o igual a  10.000. La validación debe ser hecha con la función “Is()” dentro del “main”.
*/

// Definimos una estructura llamada salaryError que tendrá un campo message de tipo string
type salaryError struct {
	message string
}

// Implementamos el método Error() para la estructura salaryError
func (salary *salaryError) Error() string {
	return salary.message
}

func main() {
	var salary int = 9000 // Creamos una variable int salary con valor 9000

	// Comprobamos si el valor de salary es menor o igual a 10000
	if salary <= 10000 {
		// Creamos un nuevo error personalizado de tipo salaryError con el mensaje indicado
		err := &salaryError{"error: el salario es menor a 10.000"}

		// Usamos la función Is() de la librería errors para comprobar si el error es de tipo salaryError
		if errors.Is(err, &salaryError{}) {
			fmt.Println(err) // Imprimimos el mensaje de error en la consola
		}
	}
}
