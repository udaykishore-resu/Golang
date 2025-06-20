package iter

import (
	"fmt"
	"iter"
)

func IterSlice() {
	nums := []int{10, 20, 30, 40, 50}

	for v := range iter.Slice(nums) {
		fmt.Println(v)
	}
}
