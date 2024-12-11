package main

import (
	"fmt"
	"sort"
)

func main() {
	//sorting with anonymous function -- highorder function
	fruits := []string{"apple", "banana", "orange", "grape"}
	sort.Slice(fruits, func(i, j int) bool {
		return len(fruits[i]) < len(fruits[j])
	})
	fmt.Println(fruits)

	//addition with anonymous function
	add := func(a, b int) int {
		return a + b
	}(1, 3)
	fmt.Println(add)

	//anonymous function capturing variable
	increment := func() {
		add++
	}
	increment()
	fmt.Println(add)

	addi := applyOperation(func(x, y int) int {
		return x + y
	}, 3, 5)
	fmt.Println(addi)

	sub := applyOperation(func(x, y int) int {
		return x - y
	}, 5, 3)
	fmt.Println(sub)

}

func applyOperation(op func(x int, y int) int, x, y int) int {
	return op(x, y)
}
