package main

import (
	"fmt"
	"sync"
)

func main() {
	var manager sync.WaitGroup

	for i := 1; i <= 3; i++ {

		manager.Add(1)

		go func(id int) {
			defer manager.Done()

			fmt.Println("Worker reporting for duty!\n", id)
		}(i)

	}
	manager.Wait()
	fmt.Println("All workers have completed their tasks. Manager signing off.")
}
