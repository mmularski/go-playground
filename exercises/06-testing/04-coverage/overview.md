# Test Coverage - Overview

## What is Test Coverage?

Test coverage measures how much of your code is executed during testing. It helps identify untested code and ensures comprehensive testing.

## Why Use Test Coverage?

- **Identify untested code** - Find code that isn't being tested
- **Improve test quality** - Ensure all code paths are tested
- **Maintain code quality** - Prevent regression bugs
- **Document behavior** - Tests serve as documentation
- **Set quality standards** - Establish minimum coverage requirements

## Basic Coverage Commands

```bash
# Run tests with coverage
go test -cover

# Generate coverage profile
go test -coverprofile=coverage.out

# View coverage in HTML
go tool cover -html=coverage.out

# View coverage in terminal
go tool cover -func=coverage.out

# Run with coverage and verbose output
go test -cover -v

# Set coverage threshold
go test -cover -covermode=count
```

## Coverage Types

### Statement Coverage
Tests that each statement is executed at least once.

```go
func ExampleFunction(x int) string {
    if x > 0 {        // Statement 1
        return "positive"  // Statement 2
    } else {
        return "negative"  // Statement 3
    }
}
```

### Branch Coverage
Tests that each branch (if/else) is executed.

```go
func ProcessData(data []int) string {
    if len(data) == 0 {           // Branch 1: empty slice
        return "empty"
    }

    sum := 0
    for _, value := range data {
        if value > 0 {             // Branch 2: positive values
            sum += value
        }
    }

    if sum > 100 {                 // Branch 3: high sum
        return "high"
    } else if sum > 50 {           // Branch 4: medium sum
        return "medium"
    } else {                       // Branch 5: low sum
        return "low"
    }
}
```

### Function Coverage
Tests that each function is called at least once.

## Coverage Analysis Example

### Function with Multiple Paths
```go
func ProcessData(data []int) string {
    if len(data) == 0 {
        return "empty"
    }

    sum := 0
    for _, value := range data {
        if value > 0 {
            sum += value
        }
    }

    if sum > 100 {
        return "high"
    } else if sum > 50 {
        return "medium"
    } else {
        return "low"
    }
}
```

### Coverage Test Cases
```go
func TestProcessDataCoverage(t *testing.T) {
    // Test empty slice (covers first branch)
    result := ProcessData([]int{})
    if result != "empty" {
        t.Errorf("Expected 'empty', got %s", result)
    }

    // Test high sum (covers positive values and high branch)
    result = ProcessData([]int{50, 60})
    if result != "high" {
        t.Errorf("Expected 'high', got %s", result)
    }

    // Test medium sum (covers medium branch)
    result = ProcessData([]int{30, 40})
    if result != "medium" {
        t.Errorf("Expected 'medium', got %s", result)
    }

    // Test low sum (covers low branch)
    result = ProcessData([]int{10, 20})
    if result != "low" {
        t.Errorf("Expected 'low', got %s", result)
    }

    // Test with negative values (covers positive check)
    result = ProcessData([]int{10, -5, 20})
    if result != "low" {
        t.Errorf("Expected 'low', got %s", result)
    }
}
```

## Coverage Best Practices

### 1. Aim for High Coverage
- **80-90%** is a good target for most projects
- **100%** is ideal but not always practical
- Focus on **critical code paths**

### 2. Use Coverage as a Guide
```bash
# Generate coverage report
go test -coverprofile=coverage.out

# View in browser
go tool cover -html=coverage.out -o coverage.html
```

### 3. Identify Uncovered Code
```bash
# View function-level coverage
go tool cover -func=coverage.out
```

### 4. Set Coverage Thresholds
```bash
# Fail if coverage is below 80%
go test -cover -covermode=count -coverpkg=./... -coverprofile=coverage.out
go tool cover -func=coverage.out | grep total | awk '{if($3 < 80.0) exit 1}'
```

## Coverage Tools and Integration

### CI/CD Integration
```yaml
# Example GitHub Actions
- name: Run tests with coverage
  run: |
    go test -coverprofile=coverage.out
    go tool cover -func=coverage.out
    go tool cover -html=coverage.out -o coverage.html
```

### Coverage Badges
- Use tools like `gocover.io` for coverage badges
- Display coverage in README files
- Track coverage trends over time

## Common Coverage Issues

### 1. Error Paths
```go
func Divide(a, b int) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("division by zero")  // Often untested
    }
    return float64(a) / float64(b), nil
}
```

### 2. Edge Cases
```go
func ProcessString(s string) string {
    if len(s) == 0 {
        return ""  // Edge case often missed
    }
    return strings.ToUpper(s)
}
```

### 3. Default Cases
```go
func GetStatus(code int) string {
    switch code {
    case 200:
        return "OK"
    case 404:
        return "Not Found"
    default:
        return "Unknown"  // Default case often untested
    }
}
```

## Improving Coverage

### 1. Add Missing Test Cases
```go
// Test error conditions
func TestDivideByZero(t *testing.T) {
    _, err := Divide(10, 0)
    if err == nil {
        t.Error("Expected error for division by zero")
    }
}

// Test edge cases
func TestProcessStringEmpty(t *testing.T) {
    result := ProcessString("")
    if result != "" {
        t.Errorf("Expected empty string, got %s", result)
    }
}
```

### 2. Use Table-Driven Tests
```go
func TestProcessDataComprehensive(t *testing.T) {
    tests := []struct {
        name     string
        input    []int
        expected string
    }{
        {"empty", []int{}, "empty"},
        {"single positive", []int{60}, "medium"},
        {"multiple positive", []int{50, 60}, "high"},
        {"with negative", []int{10, -5, 20}, "low"},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := ProcessData(tt.input)
            if result != tt.expected {
                t.Errorf("ProcessData(%v) = %s, want %s",
                    tt.input, result, tt.expected)
            }
        })
    }
}
```

## Example Program

```go
package main

import (
    "testing"
)

func ProcessData(data []int) string {
    if len(data) == 0 {
        return "empty"
    }

    sum := 0
    for _, value := range data {
        if value > 0 {
            sum += value
        }
    }

    if sum > 100 {
        return "high"
    } else if sum > 50 {
        return "medium"
    } else {
        return "low"
    }
}

func TestProcessDataCoverage(t *testing.T) {
    // Test all code paths for 100% coverage
    tests := []struct {
        name     string
        input    []int
        expected string
    }{
        {"empty", []int{}, "empty"},
        {"high", []int{50, 60}, "high"},
        {"medium", []int{30, 40}, "medium"},
        {"low", []int{10, 20}, "low"},
        {"with negative", []int{10, -5, 20}, "low"},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := ProcessData(tt.input)
            if result != tt.expected {
                t.Errorf("ProcessData(%v) = %s, want %s",
                    tt.input, result, tt.expected)
            }
        })
    }
}
```

## Next Steps

Now that you understand test coverage, try the tasks in `tasks.md`!