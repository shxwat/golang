package main

import (
	"errors"
	"fmt"
)

func validateOrder(price int) error {
	if price < 150 {
		return errors.New("minimum order value is 150")
	}
	return nil

}

func main() {
	err := validateOrder(100)

	if err != nil {
		fmt.Println("Order validation failed:", err)
		return
	}
	fmt.Println("Order placed successfully.")
}
