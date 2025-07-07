### Reference Types in Go: Pointers, Slices, Maps, Functions, and Channels

In Go, reference types hold references to underlying data structures rather than containing the data directly. When you assign a reference type to a new variable or pass it to a function, you're copying the reference (pointer), not the underlying data. This enables efficient sharing and modification of data. Key reference types include:

## 1. Pointers (*T)
A pointer stores the memory address of a value.
- **Zero Value** : *nil*
- **Key Use** : Indirectly access/modify data, avoid copying large structs.

```go
var x int = 10
p := &x     // p points to x
*p = 20     // Modifies x (now x == 20)
fmt.Println(x) // Output: 20

// Pointer to a struct
type Person struct{ Name string }
nani := &Person{Name: "Udaykishore"}
nani.Name = "Nani" // Equivalent to (*nani).Name
```
**Behavior**
- Copying a pointer creates a new reference to the same data.
- Dereferencing (**p*) accesses the underlying value.
- No pointer arithmetic (unlike C/C++).

## 2. Slices ([] T)
A slice is a dynamically sized window into an underlying array.

- **Structure**
    - Pointer to the array
    - Length (current elements)
    - Capacity (max elements without reallocation)
- **Zero Value**: *nil*
```go
arr := [5]int{1, 2, 3, 4, 5}
s := arr[1:4]   // Slice: [2,3,4] (len=3 cap=4)
s[0] = 99       // Modifies arr: [1,99,3,4,5]
```
**Behavior**
- Copies refer to the same underlying array:
    ```go
    s2 := s        // s2 shares arr with s
    s2[0] = 100    // Changes s[0] and arr[1]
    ```
- append() may create a new array if capacity is exceeded.
- Never use & with slices (they already contain a pointer).

## 3. Maps (map[K]V)
A map is a hash table storing key-value pairs.

- **Zero Value**: *nil* (unusable; must initialize with make() or literal).

- **Keys**: Must be comparable (no slices/functions).

```go
m := make(map[string]int)
m["a"] = 1
m2 := m          // m2 references the same map
m2["b"] = 2      // Also visible in m
```
**Behavior**
- Copies share the same underlying hash table.
- Safe for concurrent reads, but writes require synchronization (e.g., sync.Mutex).
- Deleting keys: delete(m, "a").

## 4. Functions
Functions are first-class types and can be assigned to variables.
- **Closures**: Functions can "capture" variables from their enclosing scope.
    ```go
    adder := func() func(int) int {
        sum := 0               // Captured variable
        return func(x int) int {
            sum += x          // sum persists across calls
            return sum
        }
    }

    f := adder()
    f(1)  // 1
    f(2)  // 3 (remembers sum)
    ```
**Behavior**
- Closures hold references to captured variables.
- Multiple closures sharing variables see the same data.

