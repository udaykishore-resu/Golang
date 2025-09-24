package examples

import (
	"fmt"
	"time"
)

// In this program thead execution is driven by time.Sleep
func Goroutines_fun() {
	fmt.Println("Goroutines_fun entry")
	// Common resource
	data := "a1b2c3d4e5"

	go print_num(data)
	go print_alpha(data)

	time.Sleep(10 * time.Second)
	fmt.Println("Goroutines_fun exit")
}

func print_num(s string) {
	for _, ch := range s {
		if ch >= '0' && ch <= '9' {
			time.Sleep(time.Second)
			fmt.Printf("%c\n", ch)
		}
	}
}

func print_alpha(s string) {
	for _, ch := range s {
		if ch >= 'a' && ch <= 'z' {
			time.Sleep(time.Second)
			fmt.Printf("%c\n", ch)
		}
	}
}
