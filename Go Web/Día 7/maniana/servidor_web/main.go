package main

import {
	"github.com/gin-gonic/gin"
}

func main() {

	router := http.NewServeMux() //multiplexor de peticiones, es un manejador de peticiones, las lleva a un recurso especifico

	router.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Pong"))
	})

	http.ListenAndServe(":8080", router) // escucha en el puerto 8080 y le pasa el multiplexor de peticiones
}
