package functiontypes

import "fmt"

func Closure_fun() {
	y := 10
	x := func() int {
		return y
	}()

	fmt.Println(x)
}
