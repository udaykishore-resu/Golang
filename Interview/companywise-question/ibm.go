package question

import (
	"fmt"
	"slices"
	"unicode"
)

func preProcess(sli []string) []string {
	var res []string
	for _, s := range sli {
		var temp []rune
		for j, ch := range s {
			if !unicode.IsLetter(ch) && !unicode.IsDigit(ch) {
				if j != 0 && len(temp) > 0 {
					temp = temp[:len(temp)-1]
				}
			} else {
				temp = append(temp, ch)
			}
		}
		res = append(res, (string(temp)))
	}
	return res
}

func SolveIbmQuestion() {
	org := []string{"0000", "1234", "6969"}
	inp := []string{"##12#234", "001#2#00", "6969", "0011#22#00"}

	results := preProcess(inp)

	fmt.Println("Results: ", results)

	for _, str := range results {
		if slices.Contains(org, str) {
			fmt.Println(str, " is present in ", org)
		} else {
			fmt.Println(str, " is not present in ", org)
		}
	}
}
