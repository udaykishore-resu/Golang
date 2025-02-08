package question

import (
	"fmt"
	"time"
)

const (
	msgcnt  = 5
	timeout = 10 * time.Second
)

func SolveHclQuestion() {
	messchan := make(chan string)
	ackchan := make(chan bool)

	go sendmsg(messchan, ackchan)
	go rcvmsg(messchan, ackchan)

	time.Sleep(timeout + time.Second)
	fmt.Println("process ended")
}

func sendmsg(messchan chan string, ackchan chan bool) {
	for i := 1; i <= msgcnt; i++ {
		msg := fmt.Sprintf("Message %d", i)

		select {
		case messchan <- msg:
			fmt.Println("Sender: sender sent msg: ", msg)
		case <-time.After(timeout):
			fmt.Println("Sender: Timeout sending message")
			return
		}

		select {
		case <-ackchan:
			fmt.Println("Sender: received ack for msg: ", msg)
		case <-time.After(timeout):
			fmt.Println("Sender: Timeout waiting for ack")
			return
		}
	}
	fmt.Println("Sender: All messages sent and acknowledged")
}

func rcvmsg(messchan chan string, ackchan chan bool) {
	for {
		select {
		case msg := <-messchan:
			fmt.Println("Receiver: Received ", msg)

			select {
			case ackchan <- true:
				fmt.Println("Receiver: Sent acknowledgment for ", msg)

			case <-time.After(timeout):
				fmt.Println("Receiver: Timeout sending Acknowledgement")
				return
			}

		case <-time.After(timeout):
			fmt.Println("Receiver: Timeout receiving message")
			return
		}
	}
}
