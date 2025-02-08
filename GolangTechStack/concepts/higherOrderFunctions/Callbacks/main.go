package main

import "fmt"

func PrintEven(nums []int, callback func(int)) {
	for _, ele := range nums {
		if ele%2 == 0 {
			callback(ele)
		}
	}
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8}

	PrintEven(numbers, func(num int) {
		fmt.Println(num)
	})
}
