package examples

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func WebScraper() {
	urls := []string{
		"https://httpbin.org/delay/1",
		"https://httpbin.org/delay/2",
		"https://httpbin.org/delay/3",
		"https://httpbin.org/delay/1",
	}

	start := time.Now()
	scrapeWebsites(urls)
	fmt.Printf("Total time: %v\n", time.Since(start))
}

func scrapeWebsites(urls []string) {
	var wg sync.WaitGroup

	for i, url := range urls {
		wg.Add(1)
		go func(id int, u string) {
			defer wg.Done()
			start := time.Now()
			resp, err := http.Get(u)

			if err != nil {
				fmt.Printf("Worker %d: Error fetching %s: %v\n", id, u, err)
				return
			}
			defer resp.Body.Close()

			fmt.Printf("Worker %d: %s - Status: %d, Time: %v\n",
				id, u, resp.StatusCode, time.Since(start))

		}(i+1, url)
	}
	wg.Wait()
	fmt.Println("All workers completed")
}
