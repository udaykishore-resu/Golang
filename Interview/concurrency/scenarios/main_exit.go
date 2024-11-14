package scenarios

import (
	"fmt"
	"time"
)

func MainExit() {
	fmt.Println("MainExit Entry")

	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("Go Routines ")
	}()
	fmt.Println("MainExit Exit")
}
