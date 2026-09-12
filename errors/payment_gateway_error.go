package main

import (
	"errors"
	"fmt"
)

var ErrLowBalance = errors.New("bank rejected: low balance")

func chargeBank(amount int) error {
	if amount > 500 {
		return ErrLowBalance
	}
	return nil
}
func processCheckout(amount int) error {
	err := chargeBank(amount)
	if err != nil {
		return fmt.Errorf("checkout failed: %w", err)
	}
	return nil
}
func main() {
	err := processCheckout(1000)

	if err != nil {
		if errors.Is(err, ErrLowBalance) {
			fmt.Println("Payment failed: insufficient balance.")
		} else {
			fmt.Println("Internal system error.")
		}
	}
}
