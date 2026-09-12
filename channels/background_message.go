package main

import (
	"fmt"
	"time"
)

func main() {
	channelChallenge := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		channelChallenge <- "The background task is running."

	}()

	msg := <-channelChallenge
	fmt.Println(msg)
}
