package aggregatedatatypes

import "fmt"

func DisplayArrays() {
	var a [5]int
	b := [5]int{1, 2}
	c := [...]int{1, 2, 3, 4, 5}

	fmt.Println("len of a: ", len(a), " val: ", a)
	fmt.Println("len of b: ", len(b), " val: ", b)
	fmt.Println("len of c: ", len(c), " val: ", c)

	//iterate using range - when we want both index and value
	for index, value := range c {
		fmt.Printf("Index: %d  Value: %d\n", index, value)
	}

	//traditional loop - when we want only index
	for i := 0; i < len(b); i++ {
		fmt.Println(b[i])
	}

	displayTwoDArrays()
	displayThreeDArrays()
}

func displayThreeDArrays() {
	threeDMatrix := [2][3][4]int{
		{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
			{9, 10, 11, 12},
		},
		{
			{13, 14, 15, 16},
			{17, 18, 19, 20},
			{21, 22, 23, 24},
		},
	}

	for i, matrix := range threeDMatrix {
		for j, row := range matrix {
			for k, val := range row {
				fmt.Printf("threeDMatrix[%d][%d][%d] = %d\n", i, j, k, val)
			}
		}
	}
}

func displayTwoDArrays() {
	twoDMatrix := [3][4]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}

	for i, row := range twoDMatrix {
		for j, value := range row {
			fmt.Printf("twoDMatrix[%d][%d] = %d\n", i, j, value)
		}
	}
}
