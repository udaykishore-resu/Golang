package main

import "fmt"

/*
- Signed: int, int8, int16, int32, int64
- Unsigned: uint, uint8 (alias for byte), uint16, uint32, uint64, uintptr
- Default integer types are int and uint, which vary depending on the architecture (32-bit or 64-bit).
*/

func main() {
	var x int = 40
	var y int8 = 50

	fmt.Println("int : ", x, "int8 : ", y)

	var pi float64 = 3.1415

	fmt.Println("float64 : ", pi)

	var com complex128 = complex(5, 7)

	fmt.Println("complex128 : ", com)

	var flag bool = true

	fmt.Println("bool : ", flag)

	var str string = "i am string"

	fmt.Println("string : ", str)

}
