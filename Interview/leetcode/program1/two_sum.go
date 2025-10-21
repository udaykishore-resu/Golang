package program1

import "fmt"

func Two_Sum() {
	arr := [8]int{2, 7, 11, 15, 6, 5, 1, 10}
	k := 16

	res := make(map[int]int)
	fmt.Println(arr)
	for i, ele := range arr {
		j, ok := res[k-ele]
		if ok {
			fmt.Println(j, i)
		} else {
			res[ele] = i
		}
	}
}
