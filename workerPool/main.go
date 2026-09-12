package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		fmt.Printf("Worker %d processing Image %d\n", id, job)
		time.Sleep(500 * time.Millisecond)
		results <- job * 2
	}
}
func main() {
	jobs := make(chan int, 50)
	results := make(chan int, 50)
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, jobs, results, &wg)
	}
	for i := 1; i <= 50; i++ {
		jobs <- i
	}
	close(jobs)
	go func() {
		wg.Wait()
		close(results)
	}()

	for res := range results {
		fmt.Printf("Main received processes image: %d\n", res)
	}
}
