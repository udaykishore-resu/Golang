package programs

import "fmt"

func FindBalancedStrStackBased() {
	var str string
	fmt.Scanln(&str)

	n := len(str)

	if n%2 != 0 {
		fmt.Println(-1)
		return
	}

	stack := []rune{}

	stack = getBalancedStr(stack, str)

	fmt.Println(stack)
}

func getBalancedStr(stack []rune, str string) []rune {
	for _, v := range str {
		if v == '{' {
			stack = append(stack, v)
		} else {
			/*if len(stack) > 0 && stack[len(stack)-1] == '{' {
				stack = stack[:len(stack)-1]

			} else {
				stack = append(stack, v)
			}*/
			stack = append(stack, v)
		}
	}
	return stack
}
