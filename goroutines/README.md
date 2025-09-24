# Concurrency in Go: A Practical Guide with Hands-On Examples
## Introduction
Concurrency is one of Go's most powerful features, enabling programs to handle multiple tasks simultaneously. Unlike traditional threading models, Go uses goroutines - lightweight threads managed by the Go runtime - that make concurrent programming more efficient and accessible.

## Starting with the Basics: Automaticity
Go's concurrency model is built around the concept of automatic scheduling. The Go runtime automatically manages goroutines, multiplexing them onto OS threads. This means you can create thousands of goroutines without the overhead typically associated with OS threads.
[example](examples/Sequential_execution.go)

## Introducing Goroutines
Goroutines are functions that run concurrently with other functions. They're incredibly lightweight (starting at 2KB stack size) compared to OS threads (typically 1-2MB).
[example](examples/goroutines.go)

## The Challenge of Asynchrony in Go

Goroutines make it easy to run tasks concurrently, but coordinating them correctly is the real challenge.  
Without proper synchronization, programs can run into several issues:

### ⚠️ Common Problems

#### 1. Race Conditions
When multiple goroutines access and modify shared data at the same time, the final result becomes unpredictable.

#### 2. Deadlocks
When goroutines wait on each other in a cycle (e.g., waiting on a channel that never receives), all progress halts.

#### 3. Unpredictable Execution Order
Since goroutines are scheduled by the Go runtime, their execution order is not deterministic, leading to unexpected outputs if not controlled.

## Synchronization with WaitGroups
sync.WaitGroup is used to wait for a collection of goroutines to finish executing.

### Real-World Example: Web Scraper
