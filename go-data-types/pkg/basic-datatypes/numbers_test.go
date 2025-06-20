package basicdatatypes

import (
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
)

func TestDisplayNumbersOutput(t *testing.T) {
	// Capture the standard output
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Execute function
	DisplayNumbers()

	// Restore stdout
	w.Close()
	os.Stdout = old

	// Read the standard output
	captured, err := io.ReadAll(r)
	if err != nil {
		t.Fatalf("Failed to read captured output: %v", err)
	}

	data := string(captured)

	tests := []struct {
		name string
		want string
	}{
		{"int8", "The value of i8: -10"},
		{"uint8", "The value of ui8: 10"},
		{"int16", "The value of i16: -20"},
		{"uint16", "The value of ui16: 20"},
		{"int32", "The value of i32: -30"},
		{"uint32", "The value of ui32: 30"},
		{"int64", "The value of i64: -40"},
		{"uint64", "The value of ui64: 40"},
		{"int", "The value of i: -50"},
		{"uint", "The value of ui: 50"},
		{"byte", "The value of b: 60"},
		{"rune", "The value of rn: ♥"},
		{"float32", "The value of f1: 123.45"},
		{"float64", "The value of f2: 1.7e+308"},
		{"complex64", "The value of c1: (10+0i)"},
		{"complex128", "The value of c2: (6+7i)"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !strings.Contains(data, tt.want) {
				t.Errorf("Output is missing expected value for %s \n Want : %q \n Got : %q \n",
					tt.name, tt.want, data)
			}
		})
	}
}

func TestNumberTypes(t *testing.T) {
	tests := []struct {
		name string
		data interface{}
		want string
	}{
		{"int8", int8(-10), "int8"},
		{"uint8", uint8(10), "uint8"},
		{"int16", int16(-20), "int16"},
		{"uint16", uint16(20), "uint16"},
		{"int32", int32(-30), "int32"},
		{"uint32", uint32(30), "uint32"},
		{"int64", int64(-40), "int64"},
		{"uint64", uint64(40), "uint64"},
		{"int", -50, "int"},
		{"uint", uint(50), "uint"},
		{"byte", byte(60), "uint8"},
		{"rune", '♥', "int32"},
		{"float32", float32(123.45), "float32"},
		{"float64", 1.7e+308, "float64"},
		{"complex64", complex64(10), "complex64"},
		{"complex128", 6 + 7i, "complex128"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			typeStr := fmt.Sprintf("%T", tt.data)
			if typeStr != tt.want {
				t.Errorf("Type mismatch for %s \n Got: %s \n Want: %s \n", tt.name, typeStr, tt.want)
			}
		})
	}
}
