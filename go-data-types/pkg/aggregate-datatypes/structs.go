package aggregatedatatypes

import (
	"fmt"
	"reflect"
)

type Book struct {
	Title  string
	Pages  int
	Author string
}

func DisplayStructs() {
	b := Book{Title: "Go with Nani", Pages: 108, Author: "Nani"}

	//Manually Access Fields
	fmt.Println(b.Title, b.Pages, b.Author)
	fmt.Println(b)

	//Reflection
	val := reflect.ValueOf(b)
	typ := reflect.TypeOf(b)

	fmt.Println("Val: ", val, "Type: ", typ)

	for i := 0; i < val.NumField(); i++ {
		fmt.Printf("%s: %v\n", typ.Field(i).Name, val.Field(i).Interface())
	}
}
