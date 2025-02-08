package constants

import (
	"fmt"
	"math"
)

const s string = "constant"

func ConstantsVar() {
	fmt.Println(s)
	const n = 500000000
	const d = 3e20 / n
	fmt.Println(d)
	fmt.Println(int64(d)) //type conversion
	fmt.Println(math.Sin(n))
}
