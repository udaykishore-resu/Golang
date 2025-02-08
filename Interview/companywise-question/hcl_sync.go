package question

import (
	"fmt"
	"sync"
	"time"
)

const (
	msgCnt  = 5
	timeout = 10 * time.Second
)

func main() {
	fmt.Println("main context started")
	msgchan := make(chan string)
	ackchan := make(chan bool)

	wg := sync.WaitGroup
	mu := sync.Mutex

	wg.Add(2)

	go sendMsg(msgchan, ackchan, &wg, &mu)
	go rcvMsg(msgchan, ackchan, &wg, &mu)

	wg.Wait()
	fmt.Println("main context ended")

}

func sendMsg(msgchan chan string, ackchan chan bool, wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()

	for i := 1; i < msgCnt; i++ {
		msg := fmt.Sprintf("message: %d", i)

		select {}
	}
}
