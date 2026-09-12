package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	candidates := []string{
		"Aman",
		"priya",
		"Rohan",
	}

	var wg sync.WaitGroup

	for _, candidate := range candidates {
		wg.Add(1)
		go func(candidate string) {
			defer wg.Done()
			fmt.Println("Evaluating answers for:", candidate)
			time.Sleep(2 * time.Second)
		}(candidate)

	}
	wg.Wait()
	fmt.Println("All evaluation complete. Final report generated!!")
}
