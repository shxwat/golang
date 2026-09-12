package main

import (
	"errors"
	"fmt"
)

var ErrUserNotFound = errors.New("user not found in db")

func fetchUserFromDB(id int) error {
	if id != 1 {
		return ErrUserNotFound
	}
	return nil
}
func getUserProfile(id int) error {
	err := fetchUserFromDB(id)
	if err != nil {
		return fmt.Errorf("profile fetch failed: %w", err)
	}
	return nil
}
func main() {
	err := getUserProfile(2)
	if err != nil {
		if errors.Is(err, ErrUserNotFound) {
			fmt.Println("404: User does not exist.")
		} else {
			fmt.Println("500: Unexpected error:", err)
		}
	}
}
