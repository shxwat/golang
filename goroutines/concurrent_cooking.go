package main

import (
	"fmt"
	"time"
)

func makeChai() {
	fmt.Println("Preparing tea...")
	time.Sleep(2 * time.Second)
	fmt.Println("Tea is ready.")
}
func makeMaggi() {
	fmt.Println("Preparing noodles...")
	time.Sleep(2 * time.Second)
	fmt.Println("Noodles are ready.")
}
func main() {
	fmt.Println("Starting both tasks.")

	go makeChai()
	go makeMaggi()
	fmt.Println("Both tasks are complete.")
	time.Sleep(3 * time.Second)
}
