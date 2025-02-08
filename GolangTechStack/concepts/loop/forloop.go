package loop

import "fmt"

func display1loop() {
	fmt.Println("***********I am 1st for loop***********")

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func display2loop() {
	fmt.Println("***********I am 2nd for loop***********")

	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println(i)
	}
}

func display3loop() {
	fmt.Println("***********I am 3rd for loop***********")

	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}
}
func DisplayforLoop() {
	display1loop()
	display2loop()
	display3loop()
}
