package main

import "fmt"

type Product struct {
	Name   string
	Price  int
	Rating float64
}

func showProduct(product Product) {
	fmt.Println("Name:", product.Name)
	fmt.Println("Price:", product.Price)
	fmt.Println("Rating:", product.Rating)
}

func main() {
	product1 := Product{
		Name:   "Macbook",
		Price:  150000,
		Rating: 4.9,
	}
	showProduct(product1)
}
