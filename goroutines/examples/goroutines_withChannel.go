package examples

import (
	"fmt"
)

func Goroutines_with_Channel() {
	fmt.Println("Goroutines_with_Channel function starting")

	// common resource
	data := "a1b2c3d45e"

	letterCh := make(chan rune)
	numberCh := make(chan rune)
	done := make(chan bool)

	// Letters Goroutine
	go func() {
		for ch := range letterCh {
			fmt.Printf("%c\n", ch)
		}
		done <- true
	}()

	// Numbers Goroutine
	go func() {
		for ch := range numberCh {
			fmt.Printf("%c\n", ch)
		}
		done <- true
	}()

	// Main goroutine controls the order
	for _, ch := range data {
		if ch >= 'a' && ch <= 'z' {
			letterCh <- ch
		} else if ch >= '0' && ch <= '9' {
			numberCh <- ch
		}
	}

	// Close channels after all work sent
	close(letterCh)
	close(numberCh)

	// Wait for both workers
	<-done
	<-done

	fmt.Println("\nGoroutines_with_Channel function ending")
}
