package main

import "fmt"

func main() {
	name := "Nani"
	fmt.Printf("My name is %s. Yes, you heard that right: %s\n", name, name)
	fmt.Printf("My name is %[1]s. Yes, you heard that right: %[1]s\n", name)

}
