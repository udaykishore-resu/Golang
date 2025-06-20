package basicdatatypes

import "fmt"

func DisplayStrings() {
	//explicite declaration
	var s1 string = "Explicit declaration with type and value"
	fmt.Println(s1)

	var s2 = "Declaration with type inference"
	fmt.Println(s2)

	s3 := "Short variable declaration"
	fmt.Println(s3)

	var s4 string
	fmt.Println(s4)

	var interpreted = "Interpreted string literal First line\nSecond line"
	fmt.Println(interpreted)

	var raw = `First line
	Second line`
	fmt.Println(raw)

	var first = "Udaykishore"
	var second = "Resu"
	var name = first + second

	fmt.Println(name)
}
