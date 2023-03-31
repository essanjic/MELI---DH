package main

import "fmt"

/*Ejercicio 1 - Impuestos de salario #1
En tu función “main”, definí una variable llamada “salary” y asignale un valor de tipo “int”.
Creá un error personalizado con un struct que implemente “Error()” con el mensaje “Error: el salario ingresado no alcanza el mínimo imponible" y lanzalo en caso de que “salary” sea menor a 150.000. De lo contrario, tendrás que imprimir por consola el mensaje “Debe pagar impuesto”
*/

// Definimos un struct que representa el error personalizado
type LowSalaryError struct{}

// Implementamos el método Error() de la interfaz error para poder personalizar el mensaje de error
func (e *LowSalaryError) Error() string {
	return "el salario ingresado no alcanza el mínimo imponible"
}

func main() {
	var salary int = 100000 // Definimos el salario como 100.000
	// if salary < 150000 {    // Comprobamos si el salario es menor a 150.000
	// 	err := LowSalaryError{}  // Creamos una instancia del error personalizado
	// 	fmt.Println(err.Error()) // Imprimimos el mensaje de error por consola
	// } else {
	// 	fmt.Println("Debe pagar impuesto") // Imprimimos un mensaje indicando que debe pagar impuestos
	// }
	err := validateError(salary)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Debe pagar impuesto")
	}
}

// Se crea la funcion validar error, se le pasa el salario y retorna un error
func validateError(salary int) error {
	if salary < 150000 {
		return fmt.Errorf("error: %w", &LowSalaryError{})
	}
	return nil
}
