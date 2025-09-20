package functiontypes

import "fmt"

func counter(a int) func() int {
	return func() int {
		a++
		return a
	}
}

func Closure_fun() {
	y := 10
	x := func() int {
		return y
	}()

	fmt.Println(x)

	a := counter(x)
	fmt.Println(a())
}
