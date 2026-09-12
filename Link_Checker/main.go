package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type Result struct {
	URL      string
	Status   string
	Duration time.Duration
}

func worker(jobs <-chan string, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()

	for url := range jobs {
		start := time.Now()

		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

		req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
		if err != nil {
			results <- Result{URL: url, Status: "DOWN (Req Error)", Duration: time.Since(start)}
			cancel()
			continue
		}
		resp, err := http.DefaultClient.Do(req)
		duration := time.Since(start)
		cancel()

		if err != nil {
			results <- Result{URL: url, Status: "DOWN (Timeout/Error)", Duration: duration}
			continue
		}
		resp.Body.Close()

		if resp.StatusCode >= 400 {
			results <- Result{
				URL:      url,
				Status:   fmt.Sprintf("DOWN (HTTP %d)", resp.StatusCode),
				Duration: duration,
			}
			continue
		}

		results <- Result{
			URL:      url,
			Status:   fmt.Sprintf("UP (HTTP %d)", resp.StatusCode),
			Duration: duration,
		}

	}
}
func main() {

	urls := []string{
		"https://google.com",
		"https://github.com",
		"https://stackoverflow.com",
		"https://golang.org",
		"https://go.dev",
		"https://www.wikipedia.org",
		"https://www.reddit.com",
		"https://www.microsoft.com",
		"https://www.apple.com",
		"https://www.amazon.com",

		"https://www.cloudflare.com",
		"https://www.mozilla.org",
		"https://www.linkedin.com",
		"https://www.netflix.com",
		"https://www.spotify.com",
		"https://www.dropbox.com",
		"https://www.notion.so",
		"https://www.figma.com",
		"https://www.docker.com",
		"https://kubernetes.io",

		"https://nodejs.org",
		"https://www.python.org",
		"https://www.rust-lang.org",
		"https://www.java.com",
		"https://react.dev",
		"https://nextjs.org",
		"https://vuejs.org",
		"https://angular.dev",
		"https://www.postgresql.org",
		"https://redis.io",

		"https://httpbin.org/status/200",
		"https://httpbin.org/status/201",
		"https://httpbin.org/status/301",
		"https://httpbin.org/status/400",
		"https://httpbin.org/status/401",
		"https://httpbin.org/status/403",
		"https://httpbin.org/status/404",
		"https://httpbin.org/status/500",
		"https://httpbin.org/status/502",
		"https://httpbin.org/status/503",

		"https://httpbin.org/delay/1",
		"https://httpbin.org/delay/1",
		"https://httpbin.org/delay/2",
		"https://httpbin.org/delay/2",
		"https://httpbin.org/delay/2",

		// 3 sec timeout
		"https://httpbin.org/delay/4",
		"https://httpbin.org/delay/5",
		"https://httpbin.org/delay/6",

		"https://example.com",
		"https://httpbin.org/get",
	}
	jobs := make(chan string)
	results := make(chan Result)
	var wg sync.WaitGroup

	numWorkers := 3

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(jobs, results, &wg)
	}
	go func() {
		for _, url := range urls {
			jobs <- url
		}
		close(jobs)
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	fmt.Println("URL Health Check Started")

	for res := range results {
		fmt.Printf("[%s] %s (Took: %v)\n", res.Status, res.URL, res.Duration)
	}
	fmt.Println("All checks completed successfully!!")
}
