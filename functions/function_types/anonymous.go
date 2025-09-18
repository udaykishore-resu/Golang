package functiontypes

import "fmt"

func Anonymous_fun() {
	func() {
		fmt.Println("I am anonymous fun")
	}()

	func(x int) {
		fmt.Println(x)
	}(50)
}
