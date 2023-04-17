package main

import (
	"encoding/json"
	"fmt"
	"os"
	//"github.com/gin-gonic/gin"
)

type Products struct {
	Id           int     `json:"id"`
	Name         string  `json:"name"`
	Quantity     int     `json:"quantity"`
	Code_value   string  `json:"code_value"`
	Is_published bool    `json:"is_published"`
	Expiration   string  `json:"expiration"`
	Price        float64 `json:"price"`
}

func main() {
	// importar archivo JSON
	file, err := os.Open("products.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// leer archivo JSON
	products := []Products{}
	if err = json.NewDecoder(file).Decode(&products); err != nil {
		panic(err)
	}

	// imprimir productos
	for _, product := range products {
		fmt.Printf("Product: %v	\n", product)
	}
}
