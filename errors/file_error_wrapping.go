package main

import (
	"errors"
	"fmt"
)

func readFile() error {
	return errors.New("file is missing")
}
func processFile() error {
	err := readFile()

	if err != nil {
		return fmt.Errorf("file not processed: %w", err)
	}
	return nil
}

func main() {
	err := processFile()
	if err != nil {
		fmt.Println("Error:", err)
	}
}
