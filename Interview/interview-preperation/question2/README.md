# Closures and Higher-Order Functions in Go

## What is a Closure?

A **closure** is a function that references variables from outside its own body. The function "closes over" these variables, allowing it to access and modify them even when executed outside their original scope.

### Key Characteristics:
- **Access outer variables**: Reference variables from their enclosing scope
- **Persistent state**: Maintain access to variables after the outer function completes
- **Private state**: Closed-over variables aren't directly accessible from outside

## Advantages of Using Closures

1. **Data Encapsulation**: Create private variables inaccessible from outside
2. **State Preservation**: Maintain state between function calls without global variables
3. **Function Factories**: Generate specialized functions with pre-configured behavior
4. **Callback Context**: Create functions that remember their context when passed as arguments
5. **Reduced Global Pollution**: Keep variables localized rather than using global scope

## Higher-Order Functions

**Higher-order functions** are functions that either:
1. Take one or more functions as arguments, or
2. Return a function as their result

Closures are often returned by higher-order functions to create specialized functions with preserved state.

## Examples

### 1. Basic Closure (Counter)

```go
package main

import "fmt"

func counter() func() int {
    count := 0 // Captured variable
    return func() int {
        count++
        return count
    }
}

func main() {
    myCounter := counter()
    fmt.Println(myCounter()) // 1
    fmt.Println(myCounter()) // 2
    fmt.Println(myCounter()) // 3
}
```

### 2. Function Factory

```go
package main

import "fmt"

func multiplier(factor int) func(int) int {
    return func(x int) int {
        return x * factor
    }
}

func main() {
    double := multiplier(2)
    triple := multiplier(3)
    
    fmt.Println(double(5))  // 10
    fmt.Println(triple(5))  // 15
}
```

### 3. Practical Use - Rate Limiter

```go
package main

import (
    "fmt"
    "time"
)

func rateLimiter(limit int, interval time.Duration) func() bool {
    tokens := limit
    lastRefill := time.Now()
    
    return func() bool {
        now := time.Now()
        if now.Sub(lastRefill) >= interval {
            tokens = limit
            lastRefill = now
        }
        if tokens > 0 {
            tokens--
            return true
        }
        return false
    }
}

func main() {
    limiter := rateLimiter(3, time.Second)
    for i := 0; i < 10; i++ {
        fmt.Printf("Request %d: ", i+1)
        if limiter() {
            fmt.Println("✅ Allowed")
        } else {
            fmt.Println("❌ Limited")
        }
        time.Sleep(300 * time.Millisecond)
    }
}
```

### 4. Higher-Order Function with Slice Operations

```go
package main

import "fmt"

func filter(numbers []int, test func(int) bool) []int {
    var result []int
    for _, n := range numbers {
        if test(n) {
            result = append(result, n)
        }
    }
    return result
}

func main() {
    numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    
    // Create closure that captures nothing but is passed as argument
    isEven := func(n int) bool { return n%2 == 0 }
    
    evens := filter(numbers, isEven)
    fmt.Println("Even numbers:", evens) // [2 4 6 8 10]
}
```

## Common Pitfalls and Solutions

### 1. Loop Variable Capture Gotcha

**Problem:**
```go
for i := 0; i < 3; i++ {
    go func() {
        fmt.Println(i) // Always prints 3
    }()
}
```

**Solution:**
```go
for i := 0; i < 3; i++ {
    i := i // Create a new variable per iteration
    go func() {
        fmt.Println(i) // Prints 0, 1, 2
    }()
}
```

### 2. Concurrent Access to Shared State

**Problem:** Multiple goroutines modifying captured variables can cause race conditions.

**Solution:** Use synchronization primitives like mutexes:

```go
func safeCounter() func() int {
    var (
        count int
        mu    sync.Mutex
    )
    return func() int {
        mu.Lock()
        defer mu.Unlock()
        count++
        return count
    }
}
```

## When to Use Closures

1. **Maintaining state** between function calls without global variables
2. **Creating function factories** that generate specialized functions
3. **Implementing decorators or middleware** that wrap existing functions
4. **Working with callbacks** that need to remember their context
5. **Implementing strategies or policies** that can be parameterized

## Key Takeaways

- Closures are functions that "remember" their creation environment
- Higher-order functions either take functions as arguments or return functions
- Together they enable powerful patterns like function factories and decorators
- They help write more modular, reusable, and maintainable code
- Be mindful of common pitfalls like loop variable capture and concurrent access

Closures and higher-order functions are fundamental concepts in Go that enable writing expressive, flexible code while maintaining clean encapsulation and reducing global state pollution.