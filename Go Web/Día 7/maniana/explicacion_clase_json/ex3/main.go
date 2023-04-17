package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {

	var err error

	//Encoder
	encoder := json.NewEncoder(os.Stdout)

	// Decoder

	decoder := json.NewDecoder(os.Stdin)

	// Se crea una estructura llamada data, será un map Es un JSON
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

	if err = encoder.Encode(data); err != nil {
		panic(err)
	}

	//Decoder
	toDecode := `{"age":30,"isMarried":true,"name":"Jhon","shopList":["milk","coffee","apple"]}`

	decoder = json.NewDecoder(strings.NewReader(toDecode))
	var decodeData map[string]any

	if err = decoder.Decode(&decodeData); err != nil {
		panic(err)
	}

	fmt.Printf("Decoded data: %v \n", decodeData)
}
