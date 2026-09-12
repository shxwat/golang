package main

import "fmt"

func main() {

	ch := make(chan string, 2)

	fmt.Println("Sender: adding items to the channel.")

	ch <- "work__1"
	ch <- "work__2"

	fmt.Println("Received:", <-ch)
	fmt.Println("Received:", <-ch)
}
