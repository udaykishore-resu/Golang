package main

import (
	"fmt"
)

type ByteSize float64

const (
	_           = iota
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func main() {
	type Mytype int
	var i Mytype = 33

	fmt.Println("i : ", i)

	fldrSz := GB
	fmt.Println("fldrSz ", fldrSz)

	var x int    // Zero value is 0
	var s string // Zero value is ""
	var p *int   // Zero value is nil

	fmt.Println("x : ", x)
	fmt.Println("s : ", s)
	fmt.Println("p : ", p)
}
