# Functions - Overview

## What are Functions?

Functions are reusable blocks of code that perform a specific task. In Go, functions are first-class citizens, meaning they can be assigned to variables, passed as arguments, and returned from other functions.

## Function Declaration

### Basic Function
```go
func functionName(parameter1 type1, parameter2 type2) returnType {
    // function body
    return value
}
```

### Function with No Parameters
```go
func greet() {
    fmt.Println("Hello!")
}
```

### Function with Parameters
```go
func greet(name string) {
    fmt.Printf("Hello, %s!\n", name)
}
```

### Function with Return Value
```go
func add(a, b int) int {
    return a + b
}
```

## Multiple Return Values

Go supports multiple return values, which is commonly used for error handling:

```go
func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, fmt.Errorf("division by zero")
    }
    return a / b, nil
}
```

## Named Return Values

You can name return values for clarity:

```go
func getPerson() (name string, age int, city string) {
    name = "John"
    age = 30
    city = "New York"
    return // naked return
}
```

## Variadic Functions

Functions that accept a variable number of arguments:

```go
func sum(numbers ...int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}

// Usage:
sum(1, 2, 3)     // 6
sum(1, 2, 3, 4, 5) // 15
```

## Function Types

Functions can be assigned to variables:

```go
type MathFunc func(int, int) int

var add MathFunc = func(a, b int) int {
    return a + b
}

var multiply MathFunc = func(a, b int) int {
    return a * b
}
```

## Anonymous Functions

Functions without names, often called immediately:

```go
result := func(a, b int) int {
    return a * b
}(5, 3) // result = 15
```

## Closures

Functions that capture variables from their outer scope:

```go
func createCounter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}

// Usage:
counter := createCounter()
fmt.Println(counter()) // 1
fmt.Println(counter()) // 2
fmt.Println(counter()) // 3
```

## Recursion

Functions that call themselves:

```go
func factorial(n int) int {
    if n <= 1 {
        return 1
    }
    return n * factorial(n-1)
}
```

## Function Best Practices

1. **Use descriptive names**: `calculateArea` instead of `calc`
2. **Keep functions small**: One function, one responsibility
3. **Use meaningful parameter names**
4. **Return errors when appropriate**
5. **Use defer for cleanup**:
```go
func processFile(filename string) error {
    file, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer file.Close() // Will be called when function returns

    // Process file...
    return nil
}
```

## Example Program

```go
package main

import "fmt"

// Basic function
func greet(name string) {
    fmt.Printf("Hello, %s!\n", name)
}

// Function with return value
func add(a, b int) int {
    return a + b
}

// Multiple return values
func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, fmt.Errorf("division by zero")
    }
    return a / b, nil
}

// Variadic function
func sum(numbers ...int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}

// Function that returns another function
func createMultiplier(factor int) func(int) int {
    return func(x int) int {
        return factor * x
    }
}

func main() {
    // Basic function call
    greet("Alice")

    // Function with return
    result := add(5, 3)
    fmt.Printf("5 + 3 = %d\n", result)

    // Multiple return values
    quotient, err := divide(10, 2)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
    } else {
        fmt.Printf("10 ÷ 2 = %d\n", quotient)
    }

    // Variadic function
    total := sum(1, 2, 3, 4, 5)
    fmt.Printf("Sum: %d\n", total)

    // Function that returns function
    double := createMultiplier(2)
    fmt.Printf("Double of 5: %d\n", double(5))
}
```

## Next Steps

Now that you understand functions, try the tasks in `tasks.md`!