package main

import "net/http"

func main() {
	var err error
	mux := http.NewServeMux() // multiplexor or router

	//Asociar una ruta con una función
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Pong"))
	})
	if err = http.ListenAndServe(":8081", mux); err != nil {
		panic(err)
	}
}
