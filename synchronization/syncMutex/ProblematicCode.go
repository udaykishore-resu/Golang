package syncmutex

import (
	"fmt"
	"sync"
)

type PAccount struct {
	balance int
}

func (a *PAccount) Deposit(amount int) {
	a.balance += amount
}

func ProblematicCode() {
	account := &PAccount{}
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
