package main

import "fmt"

func main() {
	fmt.Println("1. Main starts")

	defer fmt.Println("2. First Defer")
	defer fmt.Println("3. second Defer")
	defer fmt.Println("4. Third Defer")

	fmt.Println("5. Main Ends!!")
}
