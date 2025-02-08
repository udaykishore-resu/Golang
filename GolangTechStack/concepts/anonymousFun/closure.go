package anonymousfun

import "fmt"

func ClosureFun() {
	a := 10
	inc := func() int {
		a++
		return a
	}
	fmt.Println(inc())
	fmt.Println(inc())
}
