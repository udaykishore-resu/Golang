package programs

import (
	"fmt"
	"unicode"
)

func FindPalindrome() {
	input := "m1ad$#a*1m"
	fmt.Println(isPalindrome(input))
}

func isPalindrome(s string) bool {
	runes := []rune(s)
	fmt.Println()
	runes = clearRunes(runes)
	start, end := 0, len(runes)-1
	for start < end {
		if runes[start] != runes[end] {
			return false
		}
		start++
		end--
	}

	return true
}

func clearRunes(r []rune) []rune {
	var runes []rune
	for _, v := range r {
		if unicode.IsDigit(v) || unicode.IsLetter(v) {
			runes = append(runes, v)
		}
	}
	return runes
}
