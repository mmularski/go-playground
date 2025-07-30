# Variables and Data Types - Tasks

## Task 1: Variable Declaration
Create variables using different declaration methods.

**Requirements:**
- Declare a string variable using `var`
- Declare an integer variable using `:=`
- Declare multiple variables in one statement
- Print all variables

## Task 2: Different Data Types
Create variables of different data types and demonstrate their usage.

**Requirements:**
- Create variables for: int, float64, string, bool, byte, rune
- Demonstrate type conversion between types
- Show zero values for each type
- Print all variables with appropriate formatting

## Task 3: Type Conversion
Practice converting between different data types.

**Requirements:**
- Convert int to float64
- Convert float64 to int (truncate)
- Convert int to string using `fmt.Sprintf()`
- Convert string to int (if possible)
- Demonstrate what happens with invalid conversions

## Task 4: Constants
Create and use constants in your program.

**Requirements:**
- Define constants for mathematical values (π, e)
- Define constants for configuration values
- Use constants in calculations
- Demonstrate the difference between constants and variables

## Task 5: Variable Scope
Demonstrate variable scope in Go.

**Requirements:**
- Create variables in different scopes (package level, function level)
- Show how variables are accessible in different scopes
- Demonstrate shadowing (same name in different scopes)
- Show best practices for variable naming

## Task 6: Practical Application
Create a simple calculator that uses different variable types.

**Requirements:**
- Store user input in appropriate variable types
- Perform calculations with different numeric types
- Handle different data types appropriately
- Display results with proper formatting

---

## How to complete these tasks:

1. Navigate to the `student/` directory
2. Create a file called `main.go`
3. Implement each task in separate functions
4. Test your code by running: `go run main.go`

## Example structure for student/main.go:
```go
package main

import "fmt"

func main() {
    task1()
    task2()
    task3()
    task4()
    task5()
    task6()
}

func task1() {
    fmt.Println("=== Task 1: Variable Declaration ===")
    // Your code here
}

func task2() {
    fmt.Println("\n=== Task 2: Different Data Types ===")
    // Your code here
}

// ... continue with other tasks
```

## Testing Your Code

Run your program:
```bash
cd student/
go run main.go
```

You should see output showing the results of each task, demonstrating:
- Different variable declaration methods
- Various data types and their conversions
- Constants usage
- Variable scope examples
- A working calculator
```