package main

import (
	"errors"
	"fmt"
)

func withDrawMoney(balance int, amount int) (int, error) {
	if amount <= 0 {
		return 0, errors.New("invalid withdrawal amount")
	}
	if amount > balance {
		return 0, fmt.Errorf("insufficient balance: available balance is %d", balance)

	}
	remainingBalance := balance - amount
	return remainingBalance, nil
}

func main() {
	newBalance, err := withDrawMoney(5000, 1000)

	if err != nil {
		fmt.Println("Withdrawal failed:", err)
		return
	}
	fmt.Println("Transaction successful.\nRemaining balance:", newBalance)
}
