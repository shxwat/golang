package main

import "fmt"

func increment(num int) {
	num++
	fmt.Println("Inside function:", num)
}
func main() {
	x := 5
	fmt.Println("Before function call:", x)
	increment(x)
	fmt.Println("After function call:", x)
}
