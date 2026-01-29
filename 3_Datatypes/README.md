# Basic Types
Basic types, also known as built-in types, are the fundamental data types from which all other types are constructed. Rich set of built-in types: integers (int8-64), unsigned integers (uint8-64), floats (float32/64), complex numbers, booleans, strings, runes. Statically typed - types determined at compile time for early error detection and performance.

## Numeric Types
Go offers a wide variety of numeric types to suit different needs for precision and size.
- **Integers (Signed): Store whole numbers (positive, negative, or zero).**
    - int8: -128 to 127
    - int16: -32,768 to 32,767
    - int32: -2,147,483,648 to 2,147,483,647
    - int64: -9,223,372,036,854,775,808 to 9,223,372,036,854,775,807
    - int: This is the most common integer type. It is either 32 or 64 bits, depending on the target system's architecture (e.g., 64 bits on a 64-bit machine). You should use this unless you have a specific reason to choose a sized type.
- **Integers (Unsigned): Store only non-negative whole numbers.**
    - uint8: 0 to 255 (also known as byte)
    - uint16: 0 to 65,535
    - uint32: 0 to 4,294,967,295
    - uint64: 0 to 18,446,744,073,709,551,615
    - uint: 32 or 64 bits, depending on the architecture.
    - uintptr: An unsigned integer large enough to store the uninterpreted bits of a pointer value.

    ```go
        // Integer types
        var (
            a int     = 42    // Platform dependent (32 or 64 bit)
            b int8    = 127   // -128 to 127
            c int16   = 32767 // -32768 to 32767
            d int32   = 2147483647
            e int64   = 9223372036854775807
        )

        // Unsigned integers
        var (
            f uint    = 42    // Platform dependent
            g uint8   = 255   // 0 to 255 (byte is an alias for uint8)
            h uint16  = 65535 // 0 to 65535
            i uint32  = 4294967295
            j uint64  = 18446744073709551615
        )
    ```
- **Floating-Point Numbers: Store numbers with a decimal point.**
    - float32: 32-bit (single-precision)
    - float64: 64-bit (double-precision). This is the default and most common float type.
- **Complex Numbers: Store numbers with real and imaginary parts.**
    - complex64: float32 real and imaginary parts
    - complex128: float64 real and imaginary parts
    ```go
    // Floating point
    var (
        k float32 = 3.14159
        l float64 = 3.14159265359
    )

    // Complex numbers
    var (
        m complex64  = 1 + 2i
        n complex128 = 1.1 + 2.2i
    )
    ```