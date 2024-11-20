package bestpractices

import "fmt"

func DisplayZeroValues() {
	var i int    // Zero value is 0
	var s string // Zero value is ""
	var b bool   // Zero value is false
	var p *int   // Zero value is nil
	fmt.Println()
	fmt.Println(i, s, b, p)
}
