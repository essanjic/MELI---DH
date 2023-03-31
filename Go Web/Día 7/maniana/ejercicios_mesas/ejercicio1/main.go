package main

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

// Ejercicio 1 - Prueba de Ping
// Vamos a crear una aplicación Web con el framework Gin que tenga un endpoint /ping que al pegarle responda un texto que diga “pong”
// El endpoint deberá ser de método GET
// La respuesta de “pong” deberá ser enviada como texto, NO como JSON

// Ejercicio 2 - Manipulando el body

// Vamos a crear un endpoint llamado /saludo. Con una pequeña estructura con nombre y apellido que al pegarle deberá responder en texto “Hola + nombre + apellido”

// El endpoint deberá ser de método POST
// Se deberá usar el package JSON para resolver el ejercicio
// La respuesta deberá seguir esta estructura: “Hola Andrea Rivas”
// La estructura deberá ser como esta:
// {
// 		“nombre”: “Andrea”,
// 		“apellido”: “Rivas”
// }

type User struct {
	Name     string `json:"name"`
	Apellido string `json:"apellido"` // El nombre de la variable en el JSON y el formato que se le dará
}

func main() {

	//ejercicio1
	router := gin.Default() // Creamos el router que nos ayuda a agregar los endpoints
	// router.GET("/ping", func(c *gin.Context) {
	// 	c.String(200, "pong")
	// })

	//ejercicio2

	var user User
	jsonData := `{"nombre": "Andrea","apellido": "Rivas"}`
	json.Unmarshal([]byte(jsonData), &user)

	router.POST("/saludo", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hola " + user.Name + " " + user.Apellido,
		})
		router.Run(":8080") // Corremos el servidor en el puerto 8080

	})
}
