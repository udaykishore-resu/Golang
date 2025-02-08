package compositetypes

import "fmt"

type Person struct {
	Name string
	Age  int
	Address string
}

func add(a, b int) int {
	return a + b
}

func DisplayComposite() {
	var arr [5]int = [5]int{1, 2, 3, 4, 5} // array of 5 integers
	fmt.Println(arr)
	arr[0] = 1 
	arr[1] = 2
	arr[2] = 3
	arr[3] = 4
	arr[4] = 5
	// arr[5] = 6 // If you will try to add an element, it will panic at compile time. COMPILE TIME ERROR

	var slice []int = []int{1, 2, 3, 4, 5}
	slice = append(slice, 6) // Dynamic resizing by adding element.
	fmt.Println(slice)

	var dic map[string]int = map[string]int{"one": 1, "two": 2, "three": 3}
	fmt.Println(dic)

	var p Person = Person{Name: "Nani", Age: 33,Address: "Khammam"}
	fmt.Println(p)

	var i int = 40
	fmt.Println("i : ", i)

	var ptr *int = &i
	fmt.Println("*ptr : ", *ptr)
	*ptr = 100
	fmt.Println("i : ", i)

	var fn func(int, int) int = add
	fmt.Println("fun as a argument: ", fn(i, arr[4]))
}
