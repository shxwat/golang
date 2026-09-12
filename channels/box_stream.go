package main

import (
	"fmt"
	"time"
)

func main() {
	challengeThree := make(chan string)

	go func() {
		for i := 1; i <= 3; i++ {
			challengeThree <- fmt.Sprintf("Sending Box %d", i)
			time.Sleep(1 * time.Second)
		}
		close(challengeThree)
	}()
	for msg := range challengeThree {
		fmt.Println("Received", msg)
	}
	fmt.Println("All tasks are done. ")
}
