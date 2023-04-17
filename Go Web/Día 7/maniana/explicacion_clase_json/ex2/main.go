package main

import (
	"encoding/json"
	"os"
)

func main() {
	var err error
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
	standardEncoder := json.NewEncoder(os.Stdout)

	if err = standardEncoder.Encode(data); err != nil {
		panic(err)
	}

}
