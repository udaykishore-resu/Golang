package functiontypes

import "fmt"

func Func_With_NoArgs() {
	a := 10
	b := 20

	fmt.Println(a, b)
}

func Func_With_Args(a int, b int) {
	fmt.Println(a, b)
}
