package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalcularImpuesto(t *testing.T) {

	t.Run("Evaluar valores menores que 50000", func(t *testing.T) {
		//Arrange
		var impuesto float64
		//Act
		impuesto = ejercicioUno(30000)
		//Assert
		assert.Equal(t, 0.0, impuesto, "El impuesto calculado para un valor menor a 50.000 es incorrecto")
	})
}
