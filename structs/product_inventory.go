package main

import "fmt"

type Product struct {
	Name   string
	Price  int
	Stock  bool
	Rating float64
}

func main() {
	product1 := Product{
		Name:   "iPhone",
		Price:  120000,
		Stock:  true,
		Rating: 4.8,
	}

	product1.Price = 110000
	product1.Stock = false

	fmt.Println(product1.Name)
	fmt.Println(product1.Price)
	fmt.Println(product1.Stock)
	fmt.Println(product1.Rating)
}
