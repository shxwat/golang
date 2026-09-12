package main

import "fmt"

type Product struct {
	Name  string
	Price int
	Stock bool
}

func main() {
	product1 := Product{
		Name:  "Macbook",
		Price: 150000,
		Stock: true,
	}

	product1.Price = 200000

	fmt.Println(product1.Name)
	fmt.Println(product1.Price)
	fmt.Println(product1.Stock)
}
