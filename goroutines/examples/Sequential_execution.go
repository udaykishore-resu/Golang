package examples

import (
	"fmt"
	"time"
)

func Sequential_execution() {
	start := time.Now()

	task1()
	task2()
	task3()

	fmt.Println("The time took for sequential execution: ", time.Since(start))
}

func task3() {
	time.Sleep(time.Second)
	fmt.Println("Task3 completed")
}

func task2() {
	time.Sleep(time.Second)
	fmt.Println("Task2 completed")
}

func task1() {
	time.Sleep(time.Second)
	fmt.Println("Task1 completed")
}
