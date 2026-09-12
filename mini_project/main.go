package main

import (
	"fmt"
	"sync"
	"time"
)

type Job struct {
	ID      int
	Payload string
}

func worker(workerID int, jobs <-chan Job, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		fmt.Printf("Worker %d processing Job %d: %s\n", workerID, job.ID, job.Payload)

		time.Sleep(2 * time.Second)
		fmt.Printf("Worker %d finished Job %d\n", workerID, job.ID)
	}
	fmt.Printf("Worker %d shutting down.\n", workerID)
}

func main() {

	jobs := make(chan Job, 100)
	var wg sync.WaitGroup

	fmt.Println("Starting SaaS Background Processor...")

	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go worker(w, jobs, &wg)
	}

	for i := 1; i <= 10; i++ {
		jobs <- Job{ID: i, Payload: "Generate AI Mock Interview Data"}
	}
	close(jobs)
	wg.Wait()
	fmt.Println("All background jobs processed successfully!!...")
}
