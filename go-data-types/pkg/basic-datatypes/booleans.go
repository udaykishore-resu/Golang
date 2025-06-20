package basicdatatypes

import "fmt"

func DisplayBoolean() {
	var (
		b1 bool = true
		b2 bool = false
		b3 bool = true
		b4 bool = !b1
		b5 bool = true && false
		b6 bool = true || false
	)

	fmt.Println(b1)
	fmt.Println(b2)
	fmt.Println(b3)
	fmt.Println(b4)
	fmt.Println(b5)
	fmt.Println(b6)
}
