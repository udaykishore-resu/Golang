package examples

import (
	"fmt"
	"math/rand"
	"time"
)

func boringg(msg string, c chan string) {
	for i := 0; ; i++ {
		c <- fmt.Sprintf("%s %d\n", msg, i)
		time.Sleep(time.Duration(rand.Intn(5)) * time.Millisecond)
	}
}

func CommWithChannel() {
	c := make(chan string)

	go boringg("msg", c)
	for range 5 {
		fmt.Println(<-c)
	}
}
