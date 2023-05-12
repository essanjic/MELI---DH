package main

import (
	"net/http"

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

var data = []Saludo{
	{Nombre: "Andrea", Apellido: "Rivas"},
}

func PostSaludo(ctx *gin.Context) {
	var saludo Saludo
	ctx.BindJSON(&saludo)

	saludo = Saludo{
		Nombre:   saludo.Nombre,
		Apellido: saludo.Apellido,
	}
	data = append(data, saludo)

	ctx.IndentedJSON(http.StatusCreated, saludo)
}

func main() {
	//var err error
	router := gin.Default()
	router.POST("/saludo", ctx *gin.Context); err  {
		ctx.String(http.StatusOK, "Hola "+saludo.Nombre+" "+saludo.Apellido)
 }
	router.Run(":8084")

}
