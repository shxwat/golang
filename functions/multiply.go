package main

import "fmt"

func multiply(a, b int) int {
	return a * b
}
func main() {
	result := multiply(5, 10)
	fmt.Println(result)
}
