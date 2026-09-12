package main

import (
	"fmt"
	"sync"
)

func main() {
	files := []string{
		"movie.mp4",
		"song.mp3",
		"photo.jpg",
	}
	var wg sync.WaitGroup

	for _, file := range files {
		wg.Add(1)

		go func(file string) {
			defer wg.Done()
			fmt.Println("Downloading file:", file)

		}(file)
	}
	wg.Wait()
	fmt.Println("All files have been downloaded.")
}
