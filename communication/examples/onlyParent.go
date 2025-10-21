package examples

import (
	"fmt"
	"time"
)

func boring(msg string) {
	for i := 0; ; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Second)
	}
}

func Parent_fun() {
	go boring("boring!")
	time.Sleep(time.Second * 10)
}
