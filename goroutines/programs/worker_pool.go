package programs

import (
	"fmt"
	"time"
)

func WorkerPool() {

	const numJobs = 5
	jobs := make(chan int, numJobs)
	result := make(chan int, numJobs)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, result)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	for i := 1; i <= numJobs; i++ {
		<-result
	}
}

func worker(w int, jobs chan int, result chan int) {
	for job := range jobs {
		fmt.Println("worker ", w, "started job ", job)
		time.Sleep(time.Second)
		fmt.Println("worker ", w, "finished job ", job)
		result <- job * 2
	}
}
