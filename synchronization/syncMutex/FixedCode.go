package syncmutex

import (
	"fmt"
	"sync"
)

type FAccount struct {
	balance int
	mu      sync.Mutex
}

func (f *FAccount) Deposit(amount int) {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.balance += amount
}

func FixedCode() {
	account := &FAccount{}
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			account.Deposit(10)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(account.balance)
}
