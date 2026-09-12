package main

import (
	"errors"
	// "fmt"
	// "os"
)

// func main() {
// 	file, err := os.Open("data.txt")

// 	if err != nil {
// 		fmt.Println("Error opening file", err)
// 		return
// 	}
// 	fmt.Println("File opened succeessfully", file.Name())
// }

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}
func main() {
	divide(0, 2)
}
