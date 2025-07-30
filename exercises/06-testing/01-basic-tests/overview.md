# Basic Testing in Go - Overview

## What is Testing?

Testing is the process of verifying that your code works as expected. Go has excellent built-in support for testing through the `testing` package.

## Why Test?

- **Catch bugs early** - Find problems before they reach production
- **Refactor safely** - Make changes with confidence
- **Document behavior** - Tests show how code should work
- **Improve design** - Writing tests often reveals better code structure

## Basic Test Structure

### Test File Naming
Test files must end with `_test.go`:
```
main.go          # Your code
main_test.go     # Your tests
```

### Test Function Naming
Test functions must start with `Test`:
```go
func TestFunctionName(t *testing.T) {
    // test code
}
```

## Writing Your First Test

### Example: Testing a Calculator

```go
// calculator.go
package main

func Add(a, b int) int {
    return a + b
}

func Subtract(a, b int) int {
    return a - b
}

func Multiply(a, b int) int {
    return a * b
}

func Divide(a, b int) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("division by zero")
    }
    return float64(a) / float64(b), nil
}
```

```go
// calculator_test.go
package main

import (
    "testing"
)

func TestAdd(t *testing.T) {
    result := Add(2, 3)
    expected := 5

    if result != expected {
        t.Errorf("Add(2, 3) = %d; expected %d", result, expected)
    }
}

func TestSubtract(t *testing.T) {
    result := Subtract(5, 3)
    expected := 2

    if result != expected {
        t.Errorf("Subtract(5, 3) = %d; expected %d", result, expected)
    }
}

func TestMultiply(t *testing.T) {
    result := Multiply(4, 3)
    expected := 12

    if result != expected {
        t.Errorf("Multiply(4, 3) = %d; expected %d", result, expected)
    }
}

func TestDivide(t *testing.T) {
    // Test successful division
    result, err := Divide(10, 2)
    if err != nil {
        t.Errorf("Divide(10, 2) returned error: %v", err)
    }
    if result != 5.0 {
        t.Errorf("Divide(10, 2) = %f; expected 5.0", result)
    }

    // Test division by zero
    _, err = Divide(10, 0)
    if err == nil {
        t.Error("Divide(10, 0) should return error")
    }
}
```

## Running Tests

### Basic Commands
```bash
# Run all tests in current directory
go test

# Run tests with verbose output
go test -v

# Run tests with coverage
go test -cover

# Run tests with coverage report
go test -coverprofile=coverage.out
go tool cover -html=coverage.out
```

### Test Output
```
=== RUN   TestAdd
--- PASS: TestAdd (0.00s)
=== RUN   TestSubtract
--- PASS: TestSubtract (0.00s)
=== RUN   TestMultiply
--- PASS: TestMultiply (0.00s)
=== RUN   TestDivide
--- PASS: TestDivide (0.00s)
PASS
coverage: 100.0% of statements
ok      ./student   0.001s
```

## Testing Best Practices

### 1. Test Function Names
```go
func TestAdd(t *testing.T)           // Good
func TestAddPositiveNumbers(t *testing.T)  // Better
func TestAddWithNegativeNumbers(t *testing.T)  // Best
```

### 2. Arrange-Act-Assert Pattern
```go
func TestAdd(t *testing.T) {
    // Arrange
    a, b := 2, 3
    expected := 5

    // Act
    result := Add(a, b)

    // Assert
    if result != expected {
        t.Errorf("Add(%d, %d) = %d; expected %d", a, b, result, expected)
    }
}
```

### 3. Test Error Cases
```go
func TestDivideByZero(t *testing.T) {
    _, err := Divide(10, 0)
    if err == nil {
        t.Error("Expected error when dividing by zero")
    }
    if err.Error() != "division by zero" {
        t.Errorf("Expected 'division by zero' error, got: %v", err)
    }
}
```

### 4. Use Helper Functions
```go
func assertEqual(t *testing.T, got, want int) {
    t.Helper()
    if got != want {
        t.Errorf("got %d, want %d", got, want)
    }
}

func TestAdd(t *testing.T) {
    result := Add(2, 3)
    assertEqual(t, result, 5)
}
```

## Testing Different Types

### Testing Strings
```go
func TestGreet(t *testing.T) {
    result := Greet("Alice")
    expected := "Hello, Alice!"

    if result != expected {
        t.Errorf("Greet('Alice') = %s; expected %s", result, expected)
    }
}
```

### Testing Slices
```go
func TestSumSlice(t *testing.T) {
    numbers := []int{1, 2, 3, 4, 5}
    result := SumSlice(numbers)
    expected := 15

    if result != expected {
        t.Errorf("SumSlice(%v) = %d; expected %d", numbers, result, expected)
    }
}
```

### Testing Maps
```go
func TestCountWords(t *testing.T) {
    text := "hello world hello"
    result := CountWords(text)
    expected := map[string]int{
        "hello": 2,
        "world": 1,
    }

    if !reflect.DeepEqual(result, expected) {
        t.Errorf("CountWords('%s') = %v; expected %v", text, result, expected)
    }
}
```

## Test Organization

### Group Related Tests
```go
func TestCalculator_Add(t *testing.T) {
    tests := []struct {
        name     string
        a, b     int
        expected int
    }{
        {"positive numbers", 2, 3, 5},
        {"negative numbers", -2, -3, -5},
        {"zero", 0, 5, 5},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := Add(tt.a, tt.b)
            if result != tt.expected {
                t.Errorf("Add(%d, %d) = %d; expected %d",
                        tt.a, tt.b, result, tt.expected)
            }
        })
    }
}
```

## Example Program

```go
// main.go
package main

import "fmt"

func Greet(name string) string {
    return fmt.Sprintf("Hello, %s!", name)
}

func IsEven(n int) bool {
    return n%2 == 0
}

func Reverse(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}
```

```go
// main_test.go
package main

import "testing"

func TestGreet(t *testing.T) {
    result := Greet("Alice")
    expected := "Hello, Alice!"

    if result != expected {
        t.Errorf("Greet('Alice') = %s; expected %s", result, expected)
    }
}

func TestIsEven(t *testing.T) {
    tests := []struct {
        input    int
        expected bool
    }{
        {2, true},
        {3, false},
        {0, true},
        {-2, true},
        {-3, false},
    }

    for _, tt := range tests {
        result := IsEven(tt.input)
        if result != tt.expected {
            t.Errorf("IsEven(%d) = %t; expected %t", tt.input, result, tt.expected)
        }
    }
}

func TestReverse(t *testing.T) {
    tests := []struct {
        input    string
        expected string
    }{
        {"hello", "olleh"},
        {"world", "dlrow"},
        {"", ""},
        {"a", "a"},
        {"123", "321"},
    }

    for _, tt := range tests {
        result := Reverse(tt.input)
        if result != tt.expected {
            t.Errorf("Reverse('%s') = '%s'; expected '%s'",
                    tt.input, result, tt.expected)
        }
    }
}
```

## Next Steps

Now that you understand basic testing, try the tasks in `tasks.md`!