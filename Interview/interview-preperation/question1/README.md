# Arrays vs Slices in Go:

## Core Definitions

### Array
A **fixed-size** collection of elements of the same type:
```go
var arr [5]int // Array of 5 integers
```

### Slice
A **dynamic, flexible view** into an array with three components:
- Pointer to underlying array
- Length (number of elements)
- Capacity (maximum elements before reallocation)

```go
var s []int // Slice of integers (no fixed size)
```

## Memory Representation

### Array
- Contiguous memory block storing all elements
- Fixed size determined at compile time
- Variable holds the actual data

Example: `[3]int{10, 20, 30}`
```
Memory: [10][20][30]
```

### Slice
- Header structure containing:
  - Pointer to underlying array
  - Length (current elements)
  - Capacity (max elements)

Example: `s := []int{10, 20, 30}`
```
Slice Header:
Data → [10][20][30]
Len  = 3
Cap  = 3
```

## Creation Methods

### Arrays
```go
// Explicit size
var arr1 [3]int           // [0,0,0]
arr2 := [3]int{1, 2, 3}   // [1,2,3]

// Inferred size
arr3 := [...]int{4, 5, 6} // [4,5,6]

// Indexed initialization
arr4 := [5]int{0: 10, 3: 40} // [10,0,0,40,0]
```

### Slices
```go
// Literal
s1 := []int{1, 2, 3}

// Using make
s2 := make([]int, 3)     // len=3, cap=3 → [0,0,0]
s3 := make([]int, 3, 5)  // len=3, cap=5 → [0,0,0]

// From array
arr := [5]int{1, 2, 3, 4, 5}
s4 := arr[1:4]           // [2,3,4]

// Empty slices
var s5 []int             // nil slice
s6 := []int{}            // empty non-nil slice
```

## Element Operations

### Adding Elements
**Arrays**: ❌ Cannot add elements (fixed size)
```go
arr := [3]int{}
arr[0] = 10 // Only update existing elements
```

**Slices**: ✅ Use `append`
```go
s := []int{1, 2}
s = append(s, 3)         // [1,2,3]
s = append(s, 4, 5)      // [1,2,3,4,5]
s = append(s, s...)      // Append slice to itself
```

### Removing Elements
**Arrays**: ❌ Cannot remove elements
```go
arr := [5]int{1, 2, 3, 4, 5}
arr[2] = 0 // Logical "removal" by zeroing
```

**Slices**: ✅ Use slicing tricks
```go
s := []int{1, 2, 3, 4, 5}

// Remove element at index 2
s = append(s[:2], s[3:]...) // [1,2,4,5]

// Remove first element
s = s[1:]                   // [2,4,5]

// Remove last element
s = s[:len(s)-1]            // [2,4]

// Remove range (index 1-2)
s = append(s[:1], s[3:]...) // [2]
```

## Key Differences Summary

| Feature | Array | Slice |
|---------|-------|-------|
| **Size** | Fixed at compile time | Dynamic, can grow/shrink |
| **Type** | `[n]T` (size part of type) | `[]T` (size not part of type) |
| **Memory** | Stores all elements | Header with pointer to array |
| **Passing** | Passed by value (copied) | Passed by reference (shared) |
| **Operations** | Element updates only | Append, remove, grow |
| **Empty Value** | Zero-filled elements | `nil` or empty non-nil |
| **Use Cases** | Fixed-size data | Dynamic collections |

## Practical Recommendations

1. **Use slices** for most cases - they're Go's primary collection type
2. **Use arrays** only when you need fixed size (e.g., cryptographic hashes, fixed matrices)
3. **Prefer `make`** when you know the required capacity to avoid reallocations
4. **Be cautious** with multiple slices sharing the same array - modifications affect all slices
5. **Use `copy()`** when you need to detach a slice from its underlying array

```go
// Detaching a slice
original := []int{1, 2, 3, 4, 5}
detached := make([]int, len(original))
copy(detached, original)
// Now modifications to detached won't affect original
```