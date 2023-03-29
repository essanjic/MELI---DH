package main

import (
	//	"errors"
	"fmt"
)

// 1. Tener una estructura llamada Product con los campos ID, Name, Price, Description y Category.
type Product struct {
	ID          int
	Name        string
	Price       float64
	Description string
	Category    string
}

// 2. Tener un slice global de Product llamado Products instanciado con valores.
type Products struct {
	Products []Product
}

// 3. 2 métodos asociados a la estructura Product: Save(), GetAll(). El método Save()
// deberá tomar el sli
func (product *Product) Save(products *Products) {
	products.Products = append(products.Products, *product)
}

// El método GetAll() deberá imprimir todos los productos guardados en el slice Products.
func (product *Product) getAll(products Products) {
	for _, p := range products.Products {
		fmt.Println(p.Name)
	}
}

//  4. Una función getById() al cual se le deberá pasar un
//     INT como parámetro y retorna el producto correspondiente al parámetro pasado.
//
// func (products *Products) getById(id){}
// 5. Ejecutar al menos una vez cada método y función definido desde main().
func main() {
	products := Products{}

	product := Product{
		ID:          1,
		Name:        "Juan",
		Price:       10.00,
		Description: "Juan",
		Category:    "Juan",
	}
	product2 := Product{
		ID:          2,
		Name:        "Juan",
		Price:       10.00,
		Description: "Juan",
		Category:    "Juan",
	}

	product.Save(&products)
	product2.Save(&products)

	fmt.Println(products.Products)

	product.getAll(products)

	// fmt.Println(products.getById(3))
}
