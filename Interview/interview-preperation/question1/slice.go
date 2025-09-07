package question1

import (
	"fmt"
)

func PrintSlice() {
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//sli := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//sli1 := make([]int, 5, 10)
	sli1 := make([]int, 9)
	sli2 := arr[1:6]
	// sli1[5] = 1
	// sli1[6] = 2
	// sli1[7] = 3
	// sli1[8] = 4
	// sli1[9] = 5

	sli1[0] = 1
	sli1[1] = 2
	sli1[2] = 3
	sli1[3] = 4
	sli1[4] = 5
	sli1[5] = 6
	sli1[6] = 7
	sli1[7] = 8
	sli1[8] = 9
	sli1 = append(sli1, 10)
	sli1 = append(sli1, 11)
	sli1 = append(sli1, 12)
	sli1 = append(sli1, 13)
	sli1 = append(sli1, 14)
	sli1 = append(sli1, 15)
	sli1 = append(sli1, 16)
	sli1 = append(sli1, 17)
	sli1 = append(sli1, 18)
	sli1 = append(sli1, 19)
	sli1 = append(sli1, 20)
	sli1 = append(sli1, 21)
	sli1 = append(sli1, 22)
	sli1 = append(sli1, 23)
	sli1 = append(sli1, 24)
	sli1 = append(sli1, 25)

	sli2[3] = 10
	sli2 = append(sli2, 50)
	sli2 = append(sli2, 51)
	sli2 = append(sli2, 52)
	sli2 = append(sli2, 53)
	sli2 = append(sli2, 54)
	//fmt.Println(arr)
	//fmt.Println(sli1)
	//fmt.Println(sli1)
	fmt.Println(sli1)
	fmt.Println(sli2)
	fmt.Println(arr)
	fmt.Println(len(sli2))
	//fmt.Println(reflect.TypeOf(a))
	//fmt.Println(reflect.TypeOf(sli))

}
