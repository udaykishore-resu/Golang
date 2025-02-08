package ifelse

import "fmt"

func DisplayIfelseState() {
	if i := 109; i%2 == 0 {
		fmt.Println("Even Number")
	} else {
		fmt.Println("Odd Number")
	}
}
