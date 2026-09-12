package main

import "fmt"

type Cart struct {
	Items     []string
	TotalBill int
}

func (c *Cart) AddItem(itemName string, price int) {
	c.Items = append(c.Items, itemName)
	c.TotalBill = c.TotalBill + price
	fmt.Printf("%s added to cart. Price: %d\n", itemName, price)

}
func main() {
	myCart := Cart{
		Items:     []string{},
		TotalBill: 0,
	}
	fmt.Println("----Cart Started----")

	myCart.AddItem("Burger", 200)
	myCart.AddItem("Coke", 100)

	fmt.Println("\n----Final Bill----")

	fmt.Printf("Items in Cart: %v\n", myCart.Items)
	fmt.Printf("Total Amount to pay: %d\n", myCart.TotalBill)
}
