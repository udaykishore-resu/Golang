package functions

import "fmt"

func displayrefVals(x *int, y *int) {
	(*x)++ // (10)++
	(*y)++ // (20)++
	fmt.Println("x : ", *x)
	fmt.Println("y : ", *y)
}

func CallbyRef() {
	var a int = 10
	var b int = 20

	displayrefVals(&a, &b) //displayrefVals(1024, 1028)

	fmt.Println("a : ", a)
	fmt.Println("b : ", b)

}
