# Go's Zero-Byte Secret: Why `struct{}` Outperforms `bool` for Scalable Deduplication

> A powerful optimization insight for Go developers working on performance-critical or large-scale systems.

---

## Table of Contents

- [Overview](#overview)
- [The Difference](#the-difference)
  - [map\[string\]bool](#mapstringbool)
  - [map\[string\]struct{}](#mapstringstruct)
- [Why Empty Struct Is Optimal](#why-empty-struct-is-optimal)
- [How to Use It](#how-to-use-it)
- [Why Does This Work?](#why-does-this-work)
- [When Should You Use This?](#when-should-you-use-this)
- [Example: Deduplicating a List](#example-deduplicating-a-list)
- [Summary](#summary)

---

## Overview

In Go, when building sets or deduplication logic, developers often reach for `map[string]bool`. It works — but it's not optimal. A more idiomatic and memory-efficient approach is to use `map[string]struct{}`, leveraging Go's **zero-sized empty struct**.

---

## The Difference

### `map[string]bool`

When you use `map[string]bool`, you're storing:
- The **string key**, and
- A **1-byte boolean value** (`true` or `false`)

### `map[string]struct{}`

When you use `map[string]struct{}`, you're storing:
- The **string key**, and
- An **empty struct value**, which occupies **zero bytes**

---

## Why Empty Struct Is Optimal

The `struct{}` type:

- Has **no fields**
- Is **zero-sized** (0 bytes)
- Still satisfies **type safety** and **semantic clarity**
- Is treated by Go's runtime as a valid value with **zero memory overhead**

---

## How to Use It

```go
seen := make(map[string]struct{})

// Mark an item as seen
seen["item123"] = struct{}{}

// Check if an item has been seen
_, exists := seen["item123"]
```

---

## Why Does This Work?

- **Map Keys vs. Values:** The key in a Go map always consumes memory (since it needs to be stored and hashed). However, the *value* can be of any type — including an empty struct.

- **Go's Runtime Optimization:** The Go runtime recognizes that `struct{}` is zero-sized and **doesn't allocate space** for it in the map's internal structure. This is a deliberate design choice in Go for efficiency.

---

## When Should You Use This?

| Use Case | Recommendation |
|---|---|
| Deduplication / set tracking | ✅ `map[T]struct{}` |
| Visited nodes in a graph | ✅ `map[T]struct{}` |
| Large-scale, performance-sensitive code | ✅ `map[T]struct{}` |
| Explicit boolean logic needed | `map[T]bool` |
| Readability for beginners is prioritized | `map[T]bool` |

- **Deduplication:** Any time you need to track a set of unique items — filtering duplicates, implementing a set, or tracking visited nodes in a graph — use `map[T]struct{}` instead of `map[T]bool`.

- **Performance-Critical Code:** For code that handles large volumes of data or runs in performance-sensitive environments, this micro-optimization can have a tangible impact.

---

## Example: Deduplicating a List

```go
func uniqueStrings(input []string) []string {
    seen := make(map[string]struct{})
    result := []string{}

    for _, s := range input {
        if _, exists := seen[s]; !exists {
            seen[s] = struct{}{}
            result = append(result, s)
        }
    }

    return result
}
```

---

## Summary

| | `map[string]bool` | `map[string]struct{}` |
|---|:---:|:---:|
| Value size | 1 byte | **0 bytes** |
| Memory efficient | ❌ | ✅ |
| GC overhead | Higher | **Lower** |
| Semantic clarity for sets | ❌ | ✅ |
| Explicit boolean logic | ✅ | ❌ |

**Use `map[string]struct{}`** when you just need to track *presence* and care about performance, memory, and GC overhead.

**Stick with `map[string]bool`** if you need to explicitly express boolean logic or prioritize readability for beginners.

---

> 💡 **Pro Tip:** This pattern extends beyond strings. `map[int]struct{}`, `map[CustomType]struct{}` — anywhere you'd build a set in Go, the empty struct is your friend.