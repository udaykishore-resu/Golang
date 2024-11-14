// main will wait until the goroutine
package scenarios

import (
	"fmt"
	"sync"
)

func WaitGroupFun() {
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(c int) {
			defer wg.Done()
			fmt.Println("Go routines ", c, " finished its task.")
		}(i)
	}
	fmt.Println("waiting for all goroutines to complete the task.")
	wg.Wait()
	fmt.Println("All go routines have finished the task. ")
}
