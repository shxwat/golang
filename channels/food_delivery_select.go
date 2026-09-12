package main

import (
	"fmt"
	"time"
)

func main() {
	swiggyChan := make(chan string)
	zomatoChan := make(chan string)

	go func() {
		fmt.Println("Swiggy Channel")

		time.Sleep(2 * time.Second)
		swiggyChan <- "Swiggy order"

	}()

	go func() {
		fmt.Println("Zomato Channel")

		time.Sleep(1 * time.Second)
		zomatoChan <- "Zomato order"

	}()

	select {
	case msg1 := <-swiggyChan:
		fmt.Println("Received order from Swiggy:", msg1)
	case msg2 := <-zomatoChan:
		fmt.Println("Received order from Zomato:", msg2)
	}

}
