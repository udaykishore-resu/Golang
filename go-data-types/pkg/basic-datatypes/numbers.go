package basicdatatypes

import "fmt"

/*
int8 		- 8-bit signed integer
uint8 		- 8-bit unsigned integer
int16 		- 16-bit signed integer
uint16 		- 16-bit unsigned integer
int32 		- 32-bit signed integer
uint32 		- 32-bit unsigned integer
int64 		- 64-bit signed integer
uint64 		- 64-bit unsigned integer
byte 		- 8-bit unsigned integer
rune 		- 32-bit signed integer
int 		- platform independent signed
uint 		- platform independent unsigned
float32		- 32-bit floating-point
float64		- 64-bit floating-point
complex64 	- 64-bit complex
complex128 	- 128-bit complex
*/
func DisplayNumbers() {
	var (
		i8   int8       = -10
		ui8  uint8      = 10
		i16  int16      = -20
		ui16 uint16     = 20
		i32  int32      = -30
		ui32 uint32     = 30
		i64  int64      = -40
		ui64 uint64     = 40
		i    int        = -50
		ui   uint       = 50
		b    byte       = 60
		rn   rune       = '♥'
		f1   float32    = 123.45
		f2   float64    = 1.7e+308
		c1   complex64  = 10
		c2   complex128 = 6 + 7i
	)

	fmt.Printf("The value of i8: %v\n", i8)
	fmt.Printf("The value of ui8: %v\n", ui8)
	fmt.Printf("The value of i16: %v\n", i16)
	fmt.Printf("The value of ui16: %v\n", ui16)
	fmt.Printf("The value of i32: %v\n", i32)
	fmt.Printf("The value of ui32: %v\n", ui32)
	fmt.Printf("The value of i64: %v\n", i64)
	fmt.Printf("The value of ui64: %v\n", ui64)
	fmt.Printf("The value of i: %v\n", i)
	fmt.Printf("The value of ui: %v\n", ui)
	fmt.Printf("The value of b: %v\n", b)
	fmt.Printf("The value of rn: %v\n", string(rn))
	fmt.Printf("The value of f1: %v\n", f1)
	fmt.Printf("The value of f2: %v\n", f2)
	fmt.Printf("The value of c1: %v\n", c1)
	fmt.Printf("The value of c2: %v\n", c2)
}
