package main

import (
	"errors"
	"fmt"
)

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

func main() {
	// Ejercicio 1 - Impuestos de salario
	// Una empresa de chocolates necesita calcular el impuesto de sus empleados al momento de depositar el sueldo, para cumplir el
	// objetivo es necesario crear una función que devuelva el impuesto de un salario.
	// Teniendo en cuenta que si la persona gana más de $50.000 se le descontará un 17 % del sueldo y si gana más de $150.000
	// se le descontará además un 10 % (27% en total).
	// Vamos a hacer funciones
	fmt.Println("Ejercicio 1 \n ")
	fmt.Println("El impuesto a pagar es: ", ejercicioUno(50000))
	fmt.Println("El impuesto a pagar es: ", ejercicioUno(150000))
	fmt.Println("El impuesto a pagar es: ", ejercicioUno(200000))
	fmt.Println("")

	fmt.Println("--------//--------\n\n Ejercicio 4\n ")

	/*Los profesores de una universidad de Colombia necesitan calcular algunas estadísticas
		de calificaciones de los/as estudiantes de un curso. Requieren calcular los valores mínimo, máximo y promedio de sus calificaciones.
	Para eso, se solicita generar una función que indique qué tipo de cálculo se quiere realizar
	(mínimo, máximo o promedio) y que devuelva otra función y un mensaje (en caso que el cálculo no esté definido) que se le pueda pasar
	una cantidad N de enteros y devuelva el cálculo que se indicó en la función anterior.
	Ejemplo:
	const (
	   minimum = "minimum"
	   average = "average"
	   maximum = "maximum"
	)


	...


	minFunc, err := operation(minimum)
	averageFunc, err := operation(average)
	maxFunc, err := operation(maximum)

	...


	minValue := minFunc(2, 3, 3, 4, 10, 2, 4, 5)
	averageValue := averageFunc(2, 3, 3, 4, 1, 2, 4, 5)
	maxValue := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)


	*/
	minFunc, err := ejercicioCuatro(minimum)
	if err != nil {
		fmt.Println(err)
		return
	}
	averageFunc, err := ejercicioCuatro(average)
	if err != nil {
		fmt.Println(err)
		return
	}
	maxFunc, err := ejercicioCuatro(maximum)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("El valor mínimo es: ", minFunc(2, 3, 3, 4, 10, 2, 4, 5))
	fmt.Println("El valor promedio es: ", averageFunc(2, 3, 3, 4, 1, 2, 4, 5))
	fmt.Println("El valor máximo es: ", maxFunc(2, 3, 3, 4, 1, 2, 4, 5))

}

func ejercicioUno(sueldo float64) float64 {
	var impuesto float64
	if sueldo > 50000 {
		impuesto += sueldo * 0.17
	} else if sueldo > 150000 {
		impuesto += sueldo * 0.27
	}
	return impuesto
}

func ejercicioCuatro(notas string) (func(...int) float64, error) {
	switch notas {
	case minimum:
		return func(nums ...int) float64 {
			if len(nums) == 0 {
				return 0
			}
			min := float64(nums[0])
			for _, num := range nums {
				if float64(num) < min {
					min = float64(num)
				}
			}
			return min
		}, nil
	case average:
		return func(nums ...int) float64 {
			if len(nums) == 0 {
				return 0
			}
			sum := 0
			for _, num := range nums {
				sum += num
			}
			return float64(sum) / float64(len(nums))
		}, nil
	case maximum:
		return func(nums ...int) float64 {
			if len(nums) == 0 {
				return 0
			}
			max := float64(nums[0])
			for _, num := range nums {
				if float64(num) > max {
					max = float64(num)
				}
			}
			return max
		}, nil
	default:
		return nil, errors.New("Operación inválida")
	}
}
