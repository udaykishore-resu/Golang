package functions

import "fmt"

func disp(nums ...int) {
	fmt.Println(nums)
}
func DisplayVariadicFun() {
	disp(1)
	disp(1, 2)
	disp(1, 2, 3)
	disp([]int{1, 2, 3, 4}...) // disp(1,2,3,4)
	numbs := []int{1, 2, 3, 4, 5}
	disp(numbs...) // disp(1,2,3,4,5)
}
