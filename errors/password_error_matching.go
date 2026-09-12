package main

import (
	"errors"
	"fmt"
)

var ErrWrongPassword = errors.New("Wrong Password")

func checkDB(password string) error {
	if password != "1234" {
		return ErrWrongPassword
	}
	return nil
}
func login(password string) error {
	err := checkDB(password)
	if err != nil {
		return fmt.Errorf("login failed: %w", err)
	}
	return nil
}

func main() {
	err := login("admin")

	if err != nil {
		if errors.Is(err, ErrWrongPassword) {
			fmt.Println("Incorrect password.")
		} else {
			fmt.Println("Unexpected error:", err)
		}
	}
}
