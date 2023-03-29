package main

import (
	"errors"
	"fmt"
)

type CustomError struct {
}

func (c *CustomError) Error() string {
	return "el salario es menor a 10,000"
}

func main() {
	var salary int = 12000

	secondError := &CustomError{}

	err := validateError(salary)

	if err != nil {
		if errors.Is(err, secondError) {
			fmt.Println(err)
			return
		}
	}
	fmt.Println("Debe pagar impuesto")
}

func validateError(salary int) error {
	if salary <= 10000 {
		return fmt.Errorf("error: %w", &CustomError{})
	}
	return nil
}
