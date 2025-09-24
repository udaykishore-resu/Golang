package programs

import (
	"fmt"
)

func ShowChannelsDeadlock() {
	fmt.Println("ShowChannelsDeadlock entry")
	ch := make(chan string, 2) //change this buffered channel size to 3 to avoid deadlock
	var msg string

	ch <- "Srinivas"
	ch <- "Satya"
	ch <- "Teja"
	msg = <-ch
	fmt.Println(msg)
	msg = <-ch
	fmt.Println(msg)
	msg = <-ch
	fmt.Println(msg)
	fmt.Println("ShowChannelsDeadlock exit")
}
