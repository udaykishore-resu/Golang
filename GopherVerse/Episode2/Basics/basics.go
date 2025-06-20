package basics

import "fmt"

type Point struct {
	x, y int
}

func Basics() {
	p := &Point{5, 6}
	// p++ //invalid operation: p++ (non-numeric type *Point)
	fmt.Println("Point:", *p)

	// Instead, you can modify the values pointed to by p:
	p.x++
	p.y++
	fmt.Println("Point:", *p)
}
