package main

import "fmt"

func printHello() {
	for i := 1; i <= 5; i++ {
		fmt.Println("Hello from Goroutine -", i)
	}
}

func main() {
	go printHello()

	for i := 1; i <= 5; i++ {
		fmt.Println("Hello from main -", i)
	}
}
