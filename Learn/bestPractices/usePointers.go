package bestpractices

import "fmt"

type LargeStruct struct {
	Data [1000000]int
}

// Using a pointer to avoid copying the large struct
func modifyData(ls *LargeStruct) {
	ls.Data[0] = 42
}

func DisplayPointersUse() {
	bigStruct := &LargeStruct{}
	modifyData(bigStruct)
	fmt.Println(bigStruct.Data[0])
}
