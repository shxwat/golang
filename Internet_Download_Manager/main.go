package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	totalDownloaded := 0

	for i := 1; i <= 4; i++ {
		wg.Add(1)
		go downloadPart(i, &wg, &mu, &totalDownloaded)
	}
	wg.Wait()
	fmt.Printf("\n All 4 parts merged successfully! Download Complete. Total Size : %d MB\n", totalDownloaded)
}

func downloadPart(partID int, wg *sync.WaitGroup, mu *sync.Mutex, total *int) {
	defer wg.Done()

	for i := 1; i <= 5; i++ {
		time.Sleep(500 * time.Millisecond)
		mu.Lock()
		*total += 5
		fmt.Printf("Worker %d downloaded 5MB. Total Progress: %d MB\n", partID, *total)
		mu.Unlock()
	}

}
