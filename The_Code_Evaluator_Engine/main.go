package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func runStudentCode(ctx context.Context, studentID int, delay time.Duration) string {
	select {
	case <-time.After(delay):
		return "Passed"
	case <-ctx.Done():
		return "Failed(Infinite Loop)"
	}
}

func evaluator(workerID int, jobs <-chan int, wg *sync.WaitGroup, mu *sync.Mutex, totalPassed *int) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("worker %d evaluating Student %d...\n", workerID, job)
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)

		delay := time.Duration(job) * time.Second

		result := runStudentCode(ctx, job, delay)
		cancel()

		if result == "Passed" {
			mu.Lock()
			*totalPassed++
			mu.Unlock()
		}
		fmt.Printf("Worker %d -> Student %d Result: %s\n", workerID, job, result)
	}
}

func main() {
	jobs := make(chan int, 10)

	var wg sync.WaitGroup
	var mu sync.Mutex
	totalPassed := 0

	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go evaluator(w, jobs, &wg, &mu, &totalPassed)
	}

	for s := 1; s <= 10; s++ {
		jobs <- s
	}
	close(jobs)
	wg.Wait()
	fmt.Printf("\nAll evaluation done! Total Students Passed: %d/10\n", totalPassed)
}
