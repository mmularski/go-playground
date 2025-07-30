# Control Flow - Tasks

## Task 1: If Statements
Create and use if statements with different conditions.

**Requirements:**
- Create if statements with simple conditions
- Create if-else statements
- Create if-else if-else chains
- Use comparison operators (==, !=, <, >, <=, >=)
- Use logical operators (&&, ||, !)

## Task 2: Switch Statements
Create and use switch statements.

**Requirements:**
- Create switch statements with default cases
- Create switch statements without default
- Use switch with expressions
- Use switch with multiple values
- Demonstrate fallthrough behavior

## Task 3: For Loops
Create different types of for loops.

**Requirements:**
- Create traditional for loops with initialization, condition, and increment
- Create while-style loops using for
- Create infinite loops with break statements
- Use continue statements to skip iterations
- Demonstrate loop control

## Task 4: Range Loops
Use range to iterate over collections.

**Requirements:**
- Use range with slices
- Use range with maps
- Use range with strings
- Access both index and value in range loops
- Use range with channels (basic usage)

## Task 5: Break and Continue
Use break and continue statements effectively.

**Requirements:**
- Use break to exit loops early
- Use continue to skip iterations
- Use labeled break and continue
- Demonstrate nested loops with break/continue
- Show practical examples

## Task 6: Nested Loops
Create and work with nested loop structures.

**Requirements:**
- Create nested for loops
- Create patterns using nested loops (triangles, rectangles)
- Use nested loops with arrays/slices
- Demonstrate loop efficiency
- Show practical applications

---

## How to complete these tasks:

1. Navigate to the `student/` directory
2. Create your code in `main.go`
3. Run your code: `go run main.go`

## Example structure for student/main.go:
```go
package main

import "fmt"

func main() {
    // Task 1: If Statements
    age := 25
    if age >= 18 {
        fmt.Println("You are an adult")
    } else {
        fmt.Println("You are a minor")
    }

    // Task 2: Switch Statements
    day := "Monday"
    switch day {
    case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
        fmt.Println("It's a weekday")
    case "Saturday", "Sunday":
        fmt.Println("It's a weekend")
    default:
        fmt.Println("Invalid day")
    }

    // Task 3: For Loops
    for i := 1; i <= 5; i++ {
        fmt.Printf("Count: %d\n", i)
    }

    // Task 4: Range Loops
    fruits := []string{"apple", "banana", "orange"}
    for index, fruit := range fruits {
        fmt.Printf("Index %d: %s\n", index, fruit)
    }

    // Task 5: Break and Continue
    for i := 1; i <= 10; i++ {
        if i == 5 {
            continue // Skip 5
        }
        if i == 8 {
            break // Stop at 8
        }
        fmt.Printf("Number: %d\n", i)
    }

    // Task 6: Nested Loops
    for i := 1; i <= 3; i++ {
        for j := 1; j <= 3; j++ {
            fmt.Printf("(%d,%d) ", i, j)
        }
        fmt.Println()
    }
}
```

## Testing Your Code

Run your program:
```bash
cd student/
go run main.go
```

You should see output demonstrating all control flow concepts.