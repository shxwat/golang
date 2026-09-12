package main

import (
	"fmt"
	"sync"
	"time"
)

func majdoor(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Worker %d is working...\n", id)
	time.Sleep(2 * time.Second)
	fmt.Printf("Worker %d completed the task.\n", id)

}
func main() {
	var wg sync.WaitGroup

	wg.Add(2)
	fmt.Println("Manager: Start working.")

	go majdoor(1, &wg)
	go majdoor(2, &wg)

	fmt.Println("Manager: Waiting for all workers.")
	wg.Wait()

	fmt.Println("Manager: All workers have left. The shop is closed.")
}
