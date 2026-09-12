package main

import (
	"context"
	"fmt"
	"time"
)

func fetchPrice(ctx context.Context, airline string, delay time.Duration, results chan<- string) {
	select {
	case <-time.After(delay):
		results <- airline + ": ₹5000"
	case <-ctx.Done():
		results <- airline + ": Timeout!!"
	}
}
func main() {

	start := time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	results := make(chan string, 3)

	go fetchPrice(ctx, "Indigo", 1*time.Second, results)

	go fetchPrice(ctx, "Air India Express", 1500*time.Millisecond, results)

	go fetchPrice(ctx, "SpiceJet", 3*time.Second, results)

	for i := 1; i <= 3; i++ {
		fmt.Println(<-results)
	}
	fmt.Println("Total Time:", time.Since(start))

}
