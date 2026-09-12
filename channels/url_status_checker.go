package main

import (
	"fmt"
	"net/http"
	"time"
)

type CheckResult struct {
	URL    string
	Status string
}

func main() {
	urls := []string{
		"https://google.com",
		"https://facebook.com",
		"https://github.com",
		"https://shxwat.in",
	}

	resultsChannel := make(chan CheckResult)
	fmt.Println("Checking websites using channels...")
	start := time.Now()

	for _, url := range urls {
		go func(site string) {
			status := "UP"
			resp, err := http.Get(site)
			if err != nil {
				status = "DOWN"
			} else {
				status = status + " " + resp.Status
			}

			resultsChannel <- CheckResult{URL: site, Status: status}
		}(url)
	}
	finalReport := make(map[string]string)

	for i := 0; i < len(urls); i++ {
		res := <-resultsChannel

		finalReport[res.URL] = res.Status
	}

	fmt.Println("\n----FINAL REPORT MAP----")
	for k, v := range finalReport {
		fmt.Println(k, "->", v)
	}
	fmt.Println("____")
	fmt.Println("Total Time:", time.Since(start))
}
