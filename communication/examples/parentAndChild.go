package examples

import (
	"fmt"
	"time"
)

func Parent_and_Child() {
	go boring("boring!")
	fmt.Println("I'm listening.")
	time.Sleep(2 * time.Second)
	fmt.Println("I'm quiting.")
}
