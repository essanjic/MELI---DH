package main

import (
	"github.com/gin-gonic/gin"
)

// Ejercicio 2 - Manipulando el body

// Vamos a crear un endpoint llamado /saludo. Con una pequeña estructura con nombre y apellido que al pegarle deberá responder en texto “Hola + nombre + apellido”

// El endpoint deberá ser de método POST
// Se deberá usar el package JSON para resolver el ejercicio
// La respuesta deberá seguir esta estructura: “Hola Andrea Rivas”
// La estructura deberá ser como esta:
//
//	{
//			“nombre”: “Andrea”,
//			“apellido”: “Rivas”
//	}
type Saludo struct {
	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
}

func main() {
	router := gin.Default()
	router.POST("/saludo", func(ctx *gin.Context) {
		var saludo Saludo
		ctx.BindJSON(&saludo)
		ctx.Request.Body.Close()
	})
}
