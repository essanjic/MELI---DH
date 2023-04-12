package main

func main() {
	// 	Ejercicio 3 - A qué mes corresponde

	// Realizar una aplicación que reciba  una variable con el número del mes.
	// Según el número, imprimir el mes que corresponda en texto.
	// ¿Se te ocurre que se puede resolver de distintas maneras? ¿Cuál elegirías y por qué?
	// Ej: 7, Julio.
	// Nota: Validar que el número del mes, sea correcto.

	numeroMes := 12
	// switch numeroMes {
	// case 1:
	// 	println("Enero")
	// case 2:
	// 	println("Febrero")
	// case 3:
	// 	println("Marzo")
	// case 4:
	// 	println("Abril")
	// case 5:
	// 	println("Mayo")
	// case 6:
	// 	println("Junio")
	// case 7:
	// 	println("Julio")
	// case 8:
	// 	println("Agosto")
	// case 9:
	// 	println("Septiembre")
	// case 10:
	// 	println("Octubre")
	// case 11:
	// 	println("Noviembre")
	// case 12:
	// 	println("Diciembre")
	// default:
	// 	println("El número del mes no es correcto")
	// }

	// Aunque se puede hacer de varias formas, según las que pensé me quedo con el map, porque es más fácil de leer y de entender, adicional el tiempo de ejecución es menor, ya que no se ejecuta todo el código, sino que solo se ejecuta el código que corresponde al número del mes que se ingresa. En el switch se ejecuta todo el código, aunque no se imprima el resultado. En el map se puede hacer una validación de que el número del mes sea correcto, ya que si no existe el número del mes, no se imprime nada. Se le agrega un if para validar que el número del mes sea correcto.
	validarMes := map[int]string{
		1:  "Enero",
		2:  "Febrero",
		3:  "Marzo",
		4:  "Abril",
		5:  "Mayo",
		6:  "Junio",
		7:  "Julio",
		8:  "Agosto",
		9:  "Septiembre",
		10: "Octubre",
		11: "Noviembre",
		12: "Diciembre",
	}

	if mes, ok := validarMes[numeroMes]; ok {
		println(mes)
	} else {
		println("El número del mes no es correcto")
	}
}
