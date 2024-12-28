package main

func factorial[P ~int](num P) P {
	if num <= 1 {
		return num
	}

	return factorial[P](num-1) * num
}

func main() {
	num := 5
	factorial(num)
}
