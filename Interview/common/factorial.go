package programs

import "fmt"

func factorialn(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorialn(n-1)
}

func Factorial() {
	num := 5
	result := factorialn(num)
	fmt.Println(result)
}
