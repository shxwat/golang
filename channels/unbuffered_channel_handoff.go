package main

import "fmt"

func main() {
	ch := make(chan string)

	go func() {
		fmt.Println("Worker: starting the task.")

		ch <- "Task completed."
	}()

	msg := <-ch

	fmt.Println("Received:", msg)
}
