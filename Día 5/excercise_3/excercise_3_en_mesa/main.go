package main

import (
	"errors"
	"fmt"
)

var customError = errors.New("error: el salario es menor a 10.000")

func main() {
	var salary int = 12000

	err := validateError(salary)

	if err != nil {
		if errors.Is(err, customError) {
			fmt.Println(err)
			return
		}
	}
	fmt.Println("Debe pagar impuesto")
}

func validateError(salary int) error {
	if salary <= 10000 {
		return customError
	}
	return nil
}
