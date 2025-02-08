package programs

import "fmt"

func FindBalancedStrBalncTrack() {
	var str string
	fmt.Scanln(&str)

	n := len(str)

	if n%2 != 0 {
		fmt.Println(-1)
		return
	}

	balanced := 0
	flip := 0

	for _, v := range str {
		if v == '{' {
			balanced++
		} else {
			balanced--
		}
		if balanced < 0 {
			flip++
			balanced += 2
		}
	}

	if balanced != 0 {
		fmt.Println("String is not Balanced")
	} else {
		fmt.Println("String is Balanced")
	}

	flip += balanced / 2

	fmt.Println(flip)

}
