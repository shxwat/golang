package main

import (
	"fmt"
	"time"
)

func main() {
	limiter := make(chan time.Time, 3)

	for i := 0; i < 3; i++ {
		limiter <- time.Now()
	}

	ticker := time.NewTicker(time.Millisecond * 333)

	go func() {
		for t := range ticker.C {
			limiter <- t
		}
	}()

	reqs := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, req := range reqs {
		reqTime := <-limiter

		fmt.Printf("Request %d processed at %v\n", req, reqTime)
	}
}
