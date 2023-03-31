package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name      string   `json:"name"`      // El nombre de la variable en el JSON y el formato que se le dará
	Age       int      `json:"age"`       // El nombre de la variable en el JSON y el formato que se le dará
	IsMarried bool     `json:"isMarried"` // El nombre de la variable en el JSON y el formato que se le dará
	ShopList  []string `json:"shopList"`  // El nombre de la variable en el JSON y el formato que se le dará
}

func main() {

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

	dataAsJSON, err := json.MarshalIndent(data, "", "  ") // Esta linea convierte el map a JSON, adicional le da formato
	if err != nil {
		panic(err)
	}

	fmt.Printf("Data as map: %v \n", data)
	fmt.Printf("Data as JSON: %s \n", dataAsJSON)

	// Crear un string JSON mediante una estructura.
	user := User{
		Name:      "Jhon",
		Age:       30,
		IsMarried: true,
		ShopList: []string{
			"milk",
			"coffee",
			"apple",
		},
	}

	userAsJSON, err := json.MarshalIndent(user, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("User as JSON: %s \n", userAsJSON)

	// Crear una estructura a partir de un JSON string
	var userMapFromJSON User

	if err = json.Unmarshal(userAsJSON, &userMapFromJSON); err != nil {
		panic(err)
	}
	fmt.Printf("User from JSON: %v \n", userMapFromJSON)

	// Crear una estructura a partir de un JSON string con formato
	var userFromJSON map[string]any

	if err = json.Unmarshal(userAsJSON, &userFromJSON); err != nil {
		panic(err)
	}

	fmt.Printf("User from JSON: %v \n", userFromJSON)

}
