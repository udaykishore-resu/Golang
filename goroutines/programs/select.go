package programs

import (
	"fmt"
	"time"
)

func Display_select() {
	ch1 := make(chan string, 3)
	ch2 := make(chan string, 3)

	go func() {
		for {
			time.Sleep(time.Second * 3)
			ch1 <- "I am thread1 - for every 3 sec"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 1)
			ch2 <- "I am thread2 - for every 1 sec"
		}
	}()

	for {
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)
		case msg2 := <-ch2:
			fmt.Println(msg2)
		}
	}
}
