package examples

import "fmt"

func Goroutines_with_select() {
	fmt.Println("Goroutines_with_select started")

	data := "a1b2c3d4e5f"

	letter := make(chan rune)
	number := make(chan rune)
	done := make(chan bool, 2)

	// letter goroutine
	go func() {
		for {
			select {
			case ch, ok := <-letter:
				if !ok {
					done <- true
					return
				}
				fmt.Printf("letter: %c\n", ch)
			}
		}
	}()

	// number goroutine
	go func() {
		for ch := range number {
			select {
			case ch, ok := <-number:
				if !ok {
					done <- true
					return
				}
				fmt.Printf("number: %c\n", ch)
			}
		}
	}()

	//main goroutine
	for _, ch := range data {
		select {
		case letter <- ch:
		case number <- ch:
		default:
			if ch >= 'a' && ch <= 'z' {
				letter <- ch
			} else if ch >= '0' && ch <= '9' {
				number <- ch
			}
		}
	}

	close(letter)
	close(number)

}
