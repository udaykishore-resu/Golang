package ifelse

import "fmt"

func DisplayIfState() {
	a := 10

	if a > 0 {
		fmt.Println("Its a positive number")
	}

	if b := -3; b < 0 {
		fmt.Println("Its not a postive number")
	}
}
