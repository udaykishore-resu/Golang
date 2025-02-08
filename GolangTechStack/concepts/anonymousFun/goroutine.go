package anonymousfun

import (
	"fmt"
	"time"
)

func DisplayParent() {
	fmt.Println("I am Parent")

	go func() {
		fmt.Println("I am Child")
	}()

	time.Sleep(1 * time.Second)

}
