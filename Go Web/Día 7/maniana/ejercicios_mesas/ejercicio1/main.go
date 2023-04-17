package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ejercicio 1 - Prueba de Ping
// Vamos a crear una aplicación Web con el framework Gin que tenga un endpoint /ping que al pegarle responda un texto que diga “pong”
// El endpoint deberá ser de método GET
// La respuesta de “pong” deberá ser enviada como texto, NO como JSON

func main() {
	router := gin.Default()
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Message: pong") // en lugar de ctx.JSON se le dio el string
	})
	router.Run(":8083")
}
