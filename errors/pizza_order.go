package main

import (
	"errors"
	"fmt"
)

func orderPizza(paise int) (string, error) {
	if paise < 300 {
		return "", errors.New("insufficient funds; order cancelled")
	}
	return "Farmhouse Pizza", nil
}
func main() {
	pizza, err := orderPizza(600)
	if err != nil {
		fmt.Println("Order failed:", err)
		return
	}

	fmt.Println("Order successful. Enjoy your", pizza)
}
