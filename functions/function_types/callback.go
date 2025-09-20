package functiontypes

import "fmt"

type callback func(int, int) int

func calculate(x, y int, operation callback) int {
	return operation(x, y)
}

func Callback_fun() {
	result := 0
	add := func(a, b int) int {
		return a + b
	}

	sub := func(a, b int) int {
		return a - b
	}

	mul := func(a, b int) int {
		return a * b
	}

	result = calculate(10, 5, add)
	fmt.Println("add: ", result)

	result = calculate(10, 5, sub)
	fmt.Println("sub: ", result)

	result = calculate(10, 5, mul)
	fmt.Println("mul: ", result)

	result = calculate(10, 5, func(a, b int) int {
		return a / b
	})
	fmt.Println("div: ", result)
}
