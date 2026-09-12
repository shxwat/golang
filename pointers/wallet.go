package main

import "fmt"

type Wallet struct {
	User    string
	Balance int
}

func (w *Wallet) Deposit(amount int) {
	w.Balance = w.Balance + amount
	fmt.Printf("%s account received a deposit of %d. New balance is: %d\n", w.User, amount, w.Balance)

}
func (w *Wallet) Pay(amount int) bool {
	if w.Balance >= amount {
		w.Balance = w.Balance - amount
		return true
	}
	return false

}

func main() {
	user := Wallet{
		User:    "Shashwat",
		Balance: 0,
	}
	fmt.Printf("The current balance of %s is %d\n", user.User, user.Balance)
	user.Deposit(500)

	isOrderPlaced := user.Pay(300)

	if isOrderPlaced {
		fmt.Println("Order placed successfully.")
	} else {
		fmt.Println("Order failed due to insufficient balance.")
	}
	fmt.Printf("Finally, %s remaining balance is: %d\n", user.User, user.Balance)
}
