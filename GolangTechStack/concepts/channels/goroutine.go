package channels

import (
	"fmt"
	"time"
)

func displayRoutine(s chan int) {
	fmt.Println("Go routines context")
	s <- 5
	s <- 10
}

func Mainthread() {
	fmt.Println("Main thread starts")
	ch := make(chan int, 2)

	go displayRoutine(ch)
	time.Sleep(1 * time.Second)

	val1 := <-ch
	val2 := <-ch

	fmt.Println(val1)
	fmt.Println(val2)
	fmt.Println("Main thread ends")
}
