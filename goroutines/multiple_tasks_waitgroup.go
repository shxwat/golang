package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Worker 1: Washing the car.")
		time.Sleep(2 * time.Second)
		fmt.Println("Worker 1: Car washing completed.")
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Worker 2: Cleaning the house.")
		time.Sleep(1 * time.Second)
		fmt.Println("Worker 2: House cleaning completed.")
	}()

	fmt.Println("Manager: Waiting for all tasks to finish.")

	wg.Wait()
	fmt.Println("All tasks are complete. The shop is closed.")
}
