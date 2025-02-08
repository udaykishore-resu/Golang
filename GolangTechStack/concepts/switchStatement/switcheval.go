package switchstatement

import (
	"fmt"
	"time"
)

func DisplaySwitchEval() {

	fmt.Println("When's Monday?")
	today := time.Now().Weekday()

	switch time.Monday {
	case today:
		fmt.Println("Today.")
	case (today + 1) % 7:
		fmt.Println("Tomorrow.")
	case (today + 2) % 7:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}
