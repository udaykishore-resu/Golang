package functions

import (
	"fmt"
)

//this function will increment the value of x and y
// print the value of x and y

func displayVals(x int, y int) { //formal arguments
	x++
	y++
	fmt.Println("x : ", x)
	fmt.Println("y : ", y)
}

func CallbyVal() { //func function_name()
	var a int = 10
	var b int = 20

	displayVals(a, b) //displayVals(10, 20) //actual arguments

	fmt.Println("a : ", a)
	fmt.Println("b : ", b)
}
