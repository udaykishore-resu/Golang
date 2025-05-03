package main

import (
	"episode1/consumer"
	"episode1/models"
	"episode1/producer"
	"time"
)

func main() {
	urls := []string{
		"https://example1.com",
		"https://example2.com",
		"https://example3.com",
		"https://example4.com",
		"https://example5.com",
		"https://example6.com",
		"https://example7.com",
		"https://example8.com",
		"https://example9.com",
		"https://example10.com",
	}

	jobs := make(chan string, len(urls))
	results := make(chan models.Result, len(urls))

	producer.StartProducerPool(len(urls), jobs, results)

	for _, url := range urls {
		jobs <- url
	}

	close(jobs)

	timeout := time.After(5 * time.Second)
	consumer.Consumer(results, len(urls), timeout)
}
