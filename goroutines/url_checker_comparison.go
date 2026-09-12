package main

import (
	"fmt"
	"net/http"
	// "sync"
	"time"
)

func main() {
	urls := []string{
		"https://google.com",
		"https://facebook.com",
		"https://github.com",
		"https://shxwat.com",
	}
	fmt.Println("Checking websites sequentially...")

	start := time.Now()

	for _, url := range urls {
		resp, err := http.Get(url)

		if err != nil {
			fmt.Println(url, "is down.")
			continue
		}
		fmt.Println(url, "is up. Status:", resp.Status)

	}
	fmt.Println("---")
	fmt.Println("Total time (sequential mode):", time.Since(start))

	// urls := []string{
	// 	"https://google.com",
	// 	"https://facebook.com",
	// 	"https://github.com",
	// 	"https://shxwat.in",
	// }

	// var wg sync.WaitGroup

	// fmt.Println("Checking websites (Fast way using CONCURRENCY)")

	// start := time.Now()
	// for _, url := range urls{
	// 	wg.Add(1)

	// 	go func(site string) {
	// 		defer wg.Done()

	// 		resp, err := http.Get(site)
	// 		if err != nil{
	// 			fmt.Println(site, "is DOWNN!!")
	// 			return
	// 		}
	// 		fmt.Println(site, "is RUNNING!!!... Status:", resp.Status)
	// 	}(url)
	// }
	// wg.Wait()

	// fmt.Println("_____")
	// fmt.Println("Total Time Taken:", time.Since(start))
}
