package ifelse

import "fmt"

func DisplayNestedIfelse() {
	a := 10
	b := 30
	c := 20

	if a > b {
		fmt.Println("A is greater")
	} else if b < c {
		fmt.Println("C is greater")
	} else {
		fmt.Println("B is greater")
	}

}
