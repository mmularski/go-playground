# Control Flow - Overview

## What is Control Flow?

Control flow determines the order in which statements are executed in a program. Go provides several control structures: if statements, switch statements, and loops.

## If Statements

### Basic If Statement
```go
if condition {
    // code to execute if condition is true
}
```

### If-Else Statement
```go
if condition {
    // code if condition is true
} else {
    // code if condition is false
}
```

### If-Else If-Else Chain
```go
if condition1 {
    // code for condition1
} else if condition2 {
    // code for condition2
} else {
    // default code
}
```

### If with Initialization
```go
if variable := expression; condition {
    // use variable here
}
// variable is not accessible here
```

## Switch Statements

### Basic Switch
```go
switch variable {
case value1:
    // code for value1
case value2:
    // code for value2
default:
    // default code
}
```

### Switch with Expression
```go
switch {
case condition1:
    // code for condition1
case condition2:
    // code for condition2
default:
    // default code
}
```

### Switch with Fallthrough
```go
switch value {
case 1:
    fmt.Println("One")
    fallthrough // continues to next case
case 2:
    fmt.Println("Two")
}
```

## For Loops

### Traditional For Loop
```go
for initialization; condition; increment {
    // loop body
}
```

### While-Style Loop
```go
for condition {
    // loop body
}
```

### Infinite Loop
```go
for {
    // loop body
    if condition {
        break
    }
}
```

### Range Loop
```go
// For slices and arrays
for index, value := range slice {
    // use index and value
}

// For maps
for key, value := range map {
    // use key and value
}

// For strings (runes)
for index, char := range string {
    // use index and char
}
```

## Break and Continue

### Break Statement
```go
for i := 0; i < 10; i++ {
    if i == 5 {
        break // exits the loop
    }
}
```

### Continue Statement
```go
for i := 0; i < 10; i++ {
    if i%2 == 0 {
        continue // skips to next iteration
    }
    // this code only runs for odd numbers
}
```

## Nested Loops

```go
for i := 1; i <= 3; i++ {
    for j := 1; j <= 3; j++ {
        fmt.Printf("(%d,%d) ", i, j)
    }
    fmt.Println()
}
```

## Best Practices

1. **Use switch for multiple conditions** instead of long if-else chains
2. **Use range for iterating** over slices, arrays, and maps
3. **Use break and continue** to control loop flow
4. **Keep loops simple** and readable
5. **Use initialization in if statements** for temporary variables

## Example Program

```go
package main

import "fmt"

func main() {
    // If statement example
    age := 18
    if age >= 18 {
        fmt.Println("Adult")
    } else {
        fmt.Println("Minor")
    }

    // Switch example
    day := "Monday"
    switch day {
    case "Monday":
        fmt.Println("Start of week")
    case "Friday":
        fmt.Println("TGIF!")
    default:
        fmt.Println("Other day")
    }

    // For loop example
    for i := 0; i < 5; i++ {
        fmt.Printf("Count: %d\n", i)
    }

    // Range example
    fruits := []string{"apple", "banana", "orange"}
    for index, fruit := range fruits {
        fmt.Printf("Index %d: %s\n", index, fruit)
    }
}
```

## Next Steps

Now that you understand control flow, try the tasks in `tasks.md`!
