package main

import "fmt"

func main() {
	ch := make(chan string, 3)

	fmt.Println("Sending messages.")

	ch <- "First Message"
	ch <- "Second Message"
	ch <- "Third Message"
	ch <- "Fourth Message"

	fmt.Println("Message", <-ch)
	fmt.Println("Message", <-ch)
	fmt.Println("Message", <-ch)
	fmt.Println("Message", <-ch)
}
