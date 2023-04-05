package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

var (
	ErrCannotpen   = errors.New("cannot open file")
	ErrCannotRead  = errors.New("cannot read file")
	ErrCanNotParse = errors.New("can not parse")
)
var products []Product

func main() {
	err := obtainData()
	if err != nil {
		panic(err)
	}
	fmt.Println(products[1])
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	router.GET("/products", getAll)
	router.GET("/products/:id", getById)
	router.GET("/products/search", searchBy)
	router.Run(":8080") }

func obtainData() error {
	file, err := os.Open("products.json")
	if err != nil {
		return ErrCannotpen
	}
	defer file.Close()

	myDecoder := json.NewDecoder(file)

	if err := myDecoder.Decode(&products); err != nil {
		return ErrCannotRead
	}
	return nil
}
func getAll(c *gin.Context) {
	c.JSON(http.StatusOK, products)

}

func getById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "Cannot parse")
		return
	}
	for _, p := range products {
		if id == p.ID {
			c.JSON(http.StatusOK, p)
			return
		}
	}
}

func searchBy(c *gin.Context) {
	price, err := strconv.ParseFloat(c.Query("price"), 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Cannot parse")
		return
	}
	var productsSlice []Product
	for _, p := range products {
		if price < p.Price {
			productsSlice = append(productsSlice, p)
		}
	}
	c.JSON(http.StatusOK, productsSlice)
}
