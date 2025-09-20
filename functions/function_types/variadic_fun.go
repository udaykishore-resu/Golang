package functiontypes

import "fmt"

func sum(a ...int) (s int) {
	for _, val := range a {
		s += val
	}
	return
}

func Variadic_fun() {
	fmt.Println(sum(1))
	fmt.Println(sum(1, 2, 3))
	fmt.Println(sum(1, 2, 3, 4, 5))
	fmt.Println(sum(1, 2, 3, 4, 5, 6, 7))
}
