package main

import (
	"fmt"
)

type CustomError struct {
	msg string
}

func (c *CustomError) Error() string {
	return "el salario ingresado no alcanza el mínimo imponible"
}

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
		return fmt.Errorf("error: %w", &CustomError{})
	}
	return nil
}
