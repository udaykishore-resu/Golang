package bestpractices

import "fmt"

func typeAssertions(val interface{}) {
	switch v := val.(type) {
	case int:
		fmt.Println(v)
	case string:
		fmt.Println(v)
	default:
		fmt.Println("unknown")
	}

}

func DisplayTypeAssertions() {
	typeAssertions(42)
	typeAssertions("nani")
	typeAssertions(3.1415)
}
