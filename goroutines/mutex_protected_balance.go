package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var balance int = 0

	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			mu.Lock()
			balance = balance + 1
			mu.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println("Final Bank Balance:", balance)
}
