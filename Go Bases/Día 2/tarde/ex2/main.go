package main

import "fmt"

func main() {
	// 	Ejercicio 2 - Préstamo

	// Un banco quiere otorgar préstamos a sus clientes, pero no todos pueden acceder a los mismos. Para ello tiene ciertas reglas para saber a qué cliente se le puede otorgar. Solo le otorga préstamos a clientes cuya edad sea mayor a 22 años, se encuentren empleados y tengan más de un año de antigüedad en su trabajo. Dentro de los préstamos que otorga no les cobrará interés a los que posean un sueldo superior a $100.000.
	// Es necesario realizar una aplicación que reciba  estas variables y que imprima un mensaje de acuerdo a cada caso.
	// Tip: tu código tiene que poder imprimir al menos 3 mensajes diferentes.
	edad := 23
	empleo := true
	antiguedad := 2
	sueldo := 10001
	if edad > 22 && empleo && antiguedad > 1 {
		if sueldo > 10000 {
			println("Tu préstamo es aprobado sin intereses :)")
		} else {
			sueldo := float64(sueldo)
			interes := sueldo * .1
			fmt.Printf("El monto a pagar con intereses es de: %.2f USD \n", sueldo+interes)
		}
	} else {
		println("Tu préstamo no es aprobado, porque no cumples con los requisitos :(")
	}
}
