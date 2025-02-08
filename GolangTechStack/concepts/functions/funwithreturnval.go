package functions

import "fmt"

func displayFun(x, y int) int {
	return x + y
}

func FunwithReturnVal() {

	res := displayFun(10, 20)
	fmt.Println(res)
}
