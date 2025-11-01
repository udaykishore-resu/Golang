# Variables & Constants
**Variables** store changeable values declared with `var` or `:=` (short declaration). 

**Constants** store unchangeable values declared with `const`. 

Variables can be explicitly typed or use type inference. Constants must be `compile-time` determinable. 

Both support block declarations and package/function scope.

## var vs :=
Go provides two main ways to declare variables: using `var` and using the short declaration operator `:=`.

The `var` keyword is used for explicit variable declarations. You can use it to define a variable with or without assigning a value. 

If no value is provided, Go assigns a `default zero value` based on the **variable type**. `var` can be used both inside and outside functions.

The `:=` syntax is a __shorthand__ for declaring and initializing a variable. 

It infers the **type** from the value and can only be used **inside functions**. 

This is a quick and convenient way to create variables without explicitly mentioning their types.

Below are the default values for different Types:

|Type|Default Value|
|-----|-----|
| **Integer**|`0` |
|**Float** |`0.0` |
|**Byte**|`0`|
|**Rune**|`0`|
|**String**| `""`|
|**Bool**| `false`|
|**Map**|`nil`|
|**Channel**| `nil`|
|**Interface**| `nil`|
|**Slices**| `nil`|
|**Pointer**| `nil`|

## const and iota
Constants declared with `const` represent unchanging **compile-time values**. 

`iota` creates successive **integer constants starting from zero, resetting per const block**. 

Useful for **enumerations**, **bit flags**, and **constant sequences without manual values**.

### Declaring Constants
#### Basic Constants
```go
// Single constant
const Pi = 3.14159

// Multiple constants
const (
    StatusOK               = 200
    StatusCreated          = 201
    StatusBadRequest       = 400
    StatusUnauthorized     = 401
    StatusForbidden        = 403
    StatusNotFound         = 404
    StatusError            = 500
)

// Typed constants
const (
    Timeout time.Duration = 30 * time.Second
    MaxSize int          = 1024
    Debug   bool         = false
)

// Untyped constants
const (
    Pi    = 3.14159
    Hello = "Hello, World"
    Name = "Udaykishore Resu"
    Yes   = true
)

// Constants can be used with different compatible types
var f float64 = Pi
var greet string = Hello
var yourName string = Name

// Basic iota
const (
    Sunday = iota       // 0
    Monday              // 1
    Tuesday             // 2
    Wednesday           // 3
    Thursday            // 4
    Friday              // 5
    Saturday            // 6
)

// Bit shifting
const (
    KB = 1 << (10 * iota)   // 1 << (10 * 0) = 1
    MB                      // 1 << (10 * 1) = 1024
    GB                      // 1 << (10 * 2) = 1048576
    TB                      // 1 << (10 * 3) = 1073741824
)

// Skip values
const (
    _           = iota              // 0
    Debug Level = 1 << iota         // 1 << 1
    Info                            // 1 << 2
    Warning                         // 1 << 3
    Error                           // 1 << 4
)

// Offset and multiplier
const (
    Offset = 2 * iota + 1  // 1
    _                      // 3
    Value                  // 5
    _                      // 7
    Result                 // 9
)

const (
    // Integer constants
    MaxInt     = 9223372036854775807
    MinInt     = -9223372036854775808
    
    // Floating-point constants
    Pi         = 3.14159265359
    E          = 2.71828182845
    
    // Complex constants
    I          = 1i
    TwoI       = 2i
)

const (
    // String constants
    Prefix     = "go_"
    Suffix     = "_temp"
    EmptyStr   = ""
    
    // Raw string literals
    MultiLine  = `Line 1
                  Line 2
                  Line 3`
)

// Boolean constants
const (
    True       = true
    False      = false
    Debug      = true
)

const (
    // Constant expressions
    Hundred    = 10 * 10
    Million    = Hundred * Hundred * Hundred
    
    // String concatenation
    FullName   = "John" + " " + "Doe"
    
    // Boolean expressions
    IsValid    = true && !false
)

// Good: Clear, descriptive names
const (
    MaxConnections   = 100
    DefaultTimeout   = 30 * time.Second
    DatabaseURL     = "postgres://localhost:5432"
)

// Avoid: Unclear names
const (
    Max   = 100     // Too vague
    TO    = 30      // Too short
    URL   = "..."   // Too generic
)
```
#### Enumerated Constants
```go
type Direction int

const (
    North Direction = iota
    East
    South
    West
)

func (d Direction) String() string {
    return [...]string{"North", "East", "South", "West"}[d]
}
```
**Output:** <br>
If d = 0 → `North`<br>
If d = 1 → `East`<br>
If d = 2 → `South`<br>
If d = 3 → `West`

#### Bitmask Constants
```go
const (
    Read Permission = 1 << iota  // 1
    Write                        // 2
    Execute                      // 4
)

// Usage
func hasPermission(p Permission) bool {
    return p&Write != 0
}
```
**Output:** <br>
p&Write checks if the bit for Write (0010) is set in p. <br>
If yes → result ≠ 0 → returns `true`. <br>
If not → result = 0 → returns `false`. <br>

#### Version Constants
```go
const (
    Major = 1
    Minor = 2
    Patch = 3
    
    Version = Major*10000 + Minor*100 + Patch
    VersionString = fmt.Sprintf("%d.%d.%d", Major, Minor, Patch)
)
```
**Output:** <br>
**Version:** 10203 <br>
**VersionString:** 1.2.3 <br>
