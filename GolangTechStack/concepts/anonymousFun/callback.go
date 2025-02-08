package anonymousfun

import "fmt"

func displayNums(nums []int, callback func(int)) {
	for _, v := range nums {
		callback(v)
	}
}

func DisplayCallback() {
	nums := []int{1, 2, 3, 4, 5}
	displayNums(nums, func(num int) {
		fmt.Println(num)
	})
}
