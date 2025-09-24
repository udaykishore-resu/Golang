package examples

import (
	"fmt"
	"time"
)

func Anonymous_Goroutine() {

	go func(msg string) {
		for i := range 3 {
			fmt.Println(msg, i)
			time.Sleep(time.Second)
		}
	}("Worker")

	time.Sleep(time.Second * 2)
	fmt.Println("Done")
}
