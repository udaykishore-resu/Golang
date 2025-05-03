package producer

import (
	"episode1/models"
	"io"
	"net/http"
	"time"
)

func producer(id int, jobs <-chan string, results chan<- models.Result) {
	for url := range jobs {
		producer := http.Client{Timeout: 2 * time.Second}
		resp, err := producer.Get(url)
		if err != nil {
			results <- models.Result{ProducerID: id, URL: url, Err: err}
			continue
		}
		respBody, _ := io.ReadAll(resp.Body)
		resp.Body.Close()
		results <- models.Result{ProducerID: id, URL: url, Bytes: len(respBody), Err: nil}
	}
}

func StartProducerPool(n int, jobs <-chan string, results chan<- models.Result) {
	for i := range n {
		go producer(i, jobs, results)
	}
}
