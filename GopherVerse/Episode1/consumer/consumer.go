package consumer

import (
	"episode1/models"
	"fmt"
	"time"
)

func Consumer(results <-chan models.Result, n int, timeout <-chan time.Time) {
	for range n {
		select {
		case res := <-results:
			if res.Err != nil {
				fmt.Printf("Producer %d FAILED to fetch %s: %v\n", res.ProducerID, res.URL, res.Err)
			} else {
				fmt.Printf("Producer %d got %d bytes from %s\n", res.ProducerID, res.Bytes, res.URL)
			}
		case <-timeout:
			fmt.Println("Timeout reached while waiting for results.")
			continue
		}
	}
}
