# Communication
From the given [example](communication/examples/parent_and_child.go) the main function couild not see the output from the other goroutine. It just printed on screen.

Real conversations require communication.

## Channels

A channel in Go provides a connection between two goroutines, allowing them to coomunicate.

```go
//Declaring and initializing
var c chan int
c = make(chan int)
// or
c := make(chan int) 

// sending data to a channel
c <- 1

// receiving data from a channel
// The "arrow" indicates the direction of the data flow.

val := <-c
```

