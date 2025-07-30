# Hello World - Overview

## What is a Go Program?

Every Go program has a specific structure:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

### Key Components:

1. **`package main`** - Every executable Go program must have a `main` package
2. **`import "fmt"`** - Import the formatting package for printing
3. **`func main()`** - The entry point of the program
4. **`fmt.Println()`** - Function to print text with a new line

## Printing in Go

Go provides several ways to print output:

### 1. `fmt.Println()` - Print with new line
```go
fmt.Println("Hello")        // Prints: Hello
fmt.Println("World")        // Prints: World
// Output:
// Hello
// World
```

### 2. `fmt.Print()` - Print without new line
```go
fmt.Print("Hello ")
fmt.Print("World")
// Output: Hello World
```

### 3. `fmt.Printf()` - Formatted printing
```go
name := "Alice"
age := 25
fmt.Printf("Name: %s, Age: %d\n", name, age)
// Output: Name: Alice, Age: 25
```

## Common Format Specifiers

- `%s` - string
- `%d` - integer
- `%f` - float
- `%t` - boolean
- `%v` - default format
- `%T` - type

## Example Program

```go
package main

import "fmt"

func main() {
    // Basic printing
    fmt.Println("Welcome to Go!")

    // Variables and printing
    name := "John"
    age := 30
    fmt.Printf("Hello, %s! You are %d years old.\n", name, age)

    // Multiple prints
    fmt.Print("This is ")
    fmt.Print("on the same line.\n")

    // Using different data types
    temperature := 22.5
    isRaining := false
    fmt.Printf("Temperature: %.1f°C, Raining: %t\n", temperature, isRaining)
}
```

## Running Your Program

To run a Go program:
```bash
go run filename.go
```

## Next Steps

Now that you understand the basics, try the tasks in `tasks.md`!