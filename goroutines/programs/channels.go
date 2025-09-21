package programs

import (
	"fmt"
	"time"
)

func print_Name(name string, ch chan string) {
	fmt.Println(name, "printName entry")
	for i := 1; i <= 3; i++ {
		ch <- name
		time.Sleep(time.Second * 1)
	}
	close(ch)
	fmt.Println(name, "printName exit")
}

func ShowChannels() {
	fmt.Println("ShowChannels entry")
	ch := make(chan string)

	go func() {
		print_Name("Om Nama Shivayah:", ch)
	}()

	fmt.Println("ShowChannels I am before loop")
	for msg := range ch {
		fmt.Println(msg)
	}
	fmt.Println("ShowChannels I am after loop")
	fmt.Println("ShowChannels exit")
}
