package programs

import (
	"fmt"
)

func ShowChannelsDeadlock() {
	fmt.Println("ShowChannelsDeadlock entry")
	ch := make(chan string, 2)
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
