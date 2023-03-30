package main

import (
	"errors"
	"fmt"
)

func main() {
	salary, err := calculateSalary(1600, 100)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("salario: ", salary)
}

func calculateSalary(hours int, valuePerHour float64) (float64, error) {
	if hours < 80 {
		return 0.0, errors.New("error: el trabajador no puede haber trabajado menos de 80 hs mensuales")
	}

	salary := float64(hours) * valuePerHour

	if salary >= 150000 {
		salary *= .9
		return salary, nil
	}
	return salary, nil
}
