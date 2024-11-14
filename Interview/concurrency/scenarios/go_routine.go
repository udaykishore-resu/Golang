package scenarios

import (
	"fmt"
	"time"
)

func GoRoutine() {
	fmt.Println("GoRoutine Entry")
	go printFun()
	time.Sleep(1 * time.Second)
	fmt.Println("GoRoutine Exit")
}

func printFun() {
	fmt.Println("printFun")
}
