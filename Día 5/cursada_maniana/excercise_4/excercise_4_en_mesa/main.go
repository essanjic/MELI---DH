package main

import (
	"fmt"
)

func main() {
	var salary int = 160000

	err := validateError(salary)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Debe pagar impuesto")
}

func validateError(salary int) error {
	if salary < 150000 {
		return fmt.Errorf("error: el minimo imponible es de 150,000 y el salario es de %d", salary)
	}
	return nil
}
