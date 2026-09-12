package main

import (
	"errors"
	"fmt"
)

func fetchFromDatabase() error {
	return errors.New("database connection failed!!")
}
func getUserData() error {
	err := fetchFromDatabase()
	if err != nil {
		return fmt.Errorf("could not get user data: %w", err)
	}
	return nil
}
func main() {
	err := getUserData()
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Done.")
}
