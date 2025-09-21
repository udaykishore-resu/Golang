package programs

import (
	"fmt"
	"sync"
	"time"
)

func printName(name string) {
	fmt.Println(name, "printName entry")
	for i := 1; i <= 3; i++ {
		fmt.Println(i, name)
		time.Sleep(time.Second * 1)
	}
	fmt.Println(name, "printName exit")
}

func ShowWaitGroup() {
	fmt.Println("ShowWaitGroup entry")
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		printName("Uday")
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("ShowWaitGroup exit")
}
