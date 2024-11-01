package main

import "fmt"

func main() {
	num := 10000000
	underscore_num := 10_000_000

	fmt.Println(num == underscore_num)
}
