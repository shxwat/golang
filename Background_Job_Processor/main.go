package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	jobs := make(chan string, 10)
	var wg sync.WaitGroup

	go func() {
		for job := range jobs {
			fmt.Println("Processing interview for:", job)
			time.Sleep(3 * time.Second)
			fmt.Println("Completed:", job)
			wg.Done()
		}
	}()
	go func() {
		for i := 1; i <= 5; i++ {
			jobName := fmt.Sprintf("User_%d", i)
			wg.Add(1)
			jobs <- jobName
			fmt.Println("Accepted request for", jobName, "- HTTP 200 returned")
			time.Sleep(1 * time.Second)
		}
	}()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)

	fmt.Println("Server is running... Try pressing Ctrl+C !")

	<-sigs
	fmt.Println("\n[!] Ctrl+C detected! Shutting down gracefully..")
	fmt.Println("[!] Waiting for active jobs to finish")
	wg.Wait()
	fmt.Println("[] All jobs completed. Server stopped safely.")
}
