package question1

import "fmt"

func PrintArray() {
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} //array len
	a := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}

	//arr[10] = 11 // this will through array out of bounce
	// as array size/len is static , cannot be changed.

	// a[12] = 13 // same as above

	fmt.Println("The len of the arr:", len(arr))
	fmt.Println("The elements in an arr:", arr)
	fmt.Println("The len of the a:", len(a))
	fmt.Println("The elements in an a:", a)
}
