package main

import (
	"errors"
	"fmt"
)

var ErrOutOfStock = errors.New("Item is out of stock")

func checkInventory(item string) error {
	return ErrOutOfStock
}
func buyItem(item string) error {
	err := checkInventory(item)

	if err != nil {
		return fmt.Errorf("cart operation failed: %w", err)
	}
	return nil
}
func main() {
	err := buyItem("iphone")
	if errors.Is(err, ErrOutOfStock) {
		fmt.Println("The product is out of stock. Please try again later.")
	} else {
		fmt.Println("Unexpected error:", err)
	}
}
