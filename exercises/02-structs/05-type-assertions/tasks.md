# Type Assertions - Tasks

## Task 1: Basic Type Assertions
Create and use basic type assertions.

**Requirements:**
- Create a function that accepts `interface{}`
- Use type assertions to check for different types
- Handle both successful and failed assertions
- Demonstrate the comma ok idiom

## Task 2: Type Switches
Create type switches to handle multiple types.

**Requirements:**
- Create a function that uses type switches
- Handle at least 5 different types
- Include a default case
- Demonstrate type switch syntax

## Task 3: Interface Type Assertions
Check if values implement specific interfaces.

**Requirements:**
- Create interfaces with methods
- Create types that implement the interfaces
- Use type assertions to check interface implementation
- Handle cases where interface is not implemented

## Task 4: Safe Type Assertions
Create functions that safely handle type assertions.

**Requirements:**
- Create functions that return values and success status
- Handle type assertion failures gracefully
- Provide default values for failed assertions
- Demonstrate error handling patterns

## Task 5: Complex Type Assertions
Work with complex types and nested structures.

**Requirements:**
- Create functions that handle slices and maps
- Use type assertions with custom structs
- Handle nested interface{} values
- Demonstrate complex type checking

## Task 6: Practical Application
Create a simple data processor using type assertions.

**Requirements:**
- Create a function that processes different data types
- Use type switches to handle various inputs
- Create a data processor that can handle mixed types
- Demonstrate real-world usage

---

## How to complete these tasks:

1. Navigate to the `student/` directory
2. Create your code in `main.go`
3. Run your code: `go run main.go`

## Example structure for student/main.go:
```go
package main

import "fmt"

// Task 1: Basic type assertions
func processBasicType(v interface{}) {
    // TODO: Use type assertion to check if v is a string
    // TODO: Use type assertion to check if v is an int
    // TODO: Handle failed assertions gracefully
}

// Task 2: Type switches
func describeType(v interface{}) {
    // TODO: Use type switch to handle different types
    // TODO: Handle string, int, bool, float64, and default cases
}

// Task 3: Interface type assertions
type Stringer interface {
    String() string
}

type Person struct {
    Name string
}

func (p Person) String() string {
    return p.Name
}

func checkStringer(v interface{}) {
    // TODO: Check if v implements Stringer interface
    // TODO: Call String() method if it does
}

// Task 4: Safe type assertions
func safeString(v interface{}) (string, bool) {
    // TODO: Safely assert v to string
    // TODO: Return string value and success status
}

func safeInt(v interface{}) (int, bool) {
    // TODO: Safely assert v to int
    // TODO: Return int value and success status
}

// Task 5: Complex type assertions
func processSlice(v interface{}) {
    // TODO: Check if v is a slice of integers
    // TODO: Process the slice if it is
}

func processMap(v interface{}) {
    // TODO: Check if v is a map[string]int
    // TODO: Process the map if it is
}

// Task 6: Practical application
type DataProcessor struct{}

func (dp DataProcessor) Process(data interface{}) string {
    // TODO: Use type switch to process different data types
    // TODO: Return appropriate string representation
    return ""
}

func main() {
    // Demonstrate your type assertions here
}
```

## Testing Your Code

Run your program:
```bash
cd student/
go run main.go
```

You should see output demonstrating all type assertion concepts.