package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {

	data := map[string]any{
		"name":      "Jhon",
		"age":       30,
		"isMarried": true,
		"shopList": []string{
			"milk",
			"coffee",
			"apple",
		},
	}
	standartEncoder := json.NewEncoder(os.Stdout) // esto es una salida por consola, esto es un writter, toma un conjunto de bytes y los escribe en la salida

	if err := standartEncoder.Encode(data); err != nil { // esto escribe en la salida y si hay un error lo lanza
		panic(err)
	}

	//Decoder

	toDecode := `{"name":"Jhon","age":30,"isMarried":true,"shopList":["milk","coffee","apple"]}`

	standardDecoder := json.NewDecoder(strings.NewReader(toDecode)) // esto es un reader, toma un conjunto de bytes y los lee

	var decodedData map[string]any

	if err := standardDecoder.Decode(&decodedData); err != nil { // esto lee y si hay un error lo lanza
		panic(err)
	}

	fmt.Printf("Decoded data: %v \n", decodedData)
}
