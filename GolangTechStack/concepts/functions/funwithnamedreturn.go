package functions

import "fmt"

func displayFunval() (x int, y int) {
	x = 10
	y = 20

	return
}

func FunwithNamedReturnVal() {
	fmt.Println(displayFunval())
}
