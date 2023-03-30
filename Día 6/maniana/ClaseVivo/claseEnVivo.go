package main

import (
	"errors"
	"fmt"
)

var (
	ErrCannotDivideByZero = errors.New("No se puede dividir por cero")
)

func Divide(num1, num2 float64) (result float64, err error) {
	if num2 == 0 {
		err = errors.New("No se puede dividir por cero")
		return
	}
	result = num1 / num2
	return
}

func main() {
	var err error

	result, err := Divide(10, 2)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)

	result, err = Divide(10, 0)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
