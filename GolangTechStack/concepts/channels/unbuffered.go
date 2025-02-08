package channels

import (
	"fmt"
	"time"
)

func UnbufferedChan() {
	ch := make(chan int)

	go func(ch chan<- int, x int) {
		time.Sleep(time.Second)
		ch <- x * x
	}(ch, 3)

	done := make(chan struct{})

	go func(ch <-chan int) {
		n := <-ch
		fmt.Println(n)
		time.Sleep(time.Second)
		done <- struct{}{}
	}(ch)

	<-done
	fmt.Println("bye")
}
