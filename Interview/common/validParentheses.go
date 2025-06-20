package programs

import "fmt"

func CheckValidParentheses() {
	params := "(({[}{}]}))"
	res := isValidParenth(params)
	fmt.Println(res)
}

func isValidParenth(params string) bool {
	stack := []rune{}
	pairs := map[rune]rune{')': '(', '}': '{', ']': '['}

	for _, v := range params {
		if i, j := pairs[v]; j {
			/*if len(stack) == 0 || stack[len(stack)-1] != i {
					return false
				}
				stack = stack[:len(stack)-1]
			} else {
				stack = append(stack, v)*/
			fmt.Println("i: ", i, " j: ", j)
		}
	}
	return len(stack) == 0
}
