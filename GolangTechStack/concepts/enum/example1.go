package enum

import "fmt"

type Status struct {
	Code    int
	Message string
}

var (
	Pending    = Status{1, "Pending"}
	InProgress = Status{2, "InProgress"}
	Completed  = Status{3, "Completed"}
	Failed     = Status{4, "Failed"}
)

const (
	KB = 1 << (10 * iota)
	MB
	GB
	TB
)

const (
	FlagRead = 1 << iota
	FlagWrite
	FlagExecute
)

const (
	Zero = iota
	One
	Two
)

type TshirtSize int

const (
	Small TshirtSize = iota + 1
	Medium
	Large
)

var sizeToString = map[Status]string{
	Pending:    "Pending",
	InProgress: "InProgress",
	Completed:  "Completed",
}

func (s Status) String() string {
	if sts, ok := sizeToString[s]; ok {
		return sts
	}
	return "Unknown"
}

func PrintEnum() {
	s := Pending
	fmt.Println(s.String())
}
