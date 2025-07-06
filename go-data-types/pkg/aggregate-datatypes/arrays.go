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
}
