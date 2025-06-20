package main

import (
	"fmt"
	"math"
	"time"
)

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	sqrt := int(math.Sqrt(float64(n)))
	for i := 2; i <= sqrt; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	start := time.Now()
	count := 0
	for i := 2; i <= 10000000; i++ {
		if isPrime(i) {
			count++
		}
	}
	elapsed := time.Since(start).Milliseconds()
	fmt.Printf("Total primes: %d\n", count)
	fmt.Printf("Execution time: %d ms\n", elapsed)
}
