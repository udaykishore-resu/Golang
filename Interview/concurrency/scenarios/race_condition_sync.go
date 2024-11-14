package scenarios

import (
	"fmt"
	"sync"
)

func RaceConditionsWait() {
	counter := 0
	var wg sync.WaitGroup
	var mt sync.Mutex

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mt.Lock()
			counter++
			mt.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println("counter :", counter)

}
