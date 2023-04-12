package main

func main() {
	// 	Ejercicio 4 - Qué edad tiene...
	// Un empleado de una empresa quiere saber el nombre y edad de uno de sus empleados. Según el siguiente mapa, ayudá a imprimir la edad de Benjamin.

	//   var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

	// Por otro lado también es necesario:
	// Saber cuántos de sus empleados son mayores de 21 años.
	// Agregar un empleado nuevo a la lista, llamado Federico que tiene 25 años.
	// Eliminar a Pedro del mapa.

	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

	// Para imprimir la edad de Benjamin
	println(employees["Benjamin"])

	// Para saber cuántos de sus empleados son mayores de 21 años
	var mayoresDe21 int
	for _, edad := range employees {
		if edad > 21 {
			mayoresDe21++
		}
	}
	println("La cantidad de empleados mayores de 21 años es:", mayoresDe21)
}
