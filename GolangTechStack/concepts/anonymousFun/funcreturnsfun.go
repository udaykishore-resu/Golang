package anonymousfun

import "fmt"

func dispReturn() func(string) string { // disp  func()
	return func(s string) string {
		return s
	}
}

func DisplayFunReturn() {

	res := dispReturn()
	fmt.Println(res("mutal"))
}
