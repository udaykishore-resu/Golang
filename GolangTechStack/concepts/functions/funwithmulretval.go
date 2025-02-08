package functions

import "fmt"

func displayFunret() (int, int) {
	a := 20
	b := 30

	return a, b
}
func FunwithMulReturnVal() {
	fmt.Println(displayFunret())
}
