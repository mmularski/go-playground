# Table-Driven Tests - Overview

## What are Table-Driven Tests?

Table-driven tests are a pattern where you define test cases in a data structure (usually a slice of structs) and iterate over them. This approach makes tests more maintainable and reduces code duplication.

## Why Use Table-Driven Tests?

- **Reduces duplication** - Write test logic once, run with multiple inputs
- **Easy to add cases** - Just add a new entry to the test data
- **Clear test coverage** - All test cases are visible in one place
- **Consistent structure** - All tests follow the same pattern

## Basic Table-Driven Test Structure

```go
func TestFunction(t *testing.T) {
    tests := []struct {
        name     string
        input    string
        expected string
    }{
        {"empty string", "", ""},
        {"single word", "hello", "HELLO"},
        {"multiple words", "hello world", "HELLO WORLD"},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := Function(tt.input)
            if result != tt.expected {
                t.Errorf("Function(%q) = %q, want %q",
                    tt.input, result, tt.expected)
            }
        })
    }
}
```

## Test Case Struct Design

### Simple Cases
```go
type testCase struct {
    input    string
    expected string
}

func TestSimpleFunction(t *testing.T) {
    tests := []testCase{
        {"hello", "HELLO"},
        {"world", "WORLD"},
        {"", ""},
    }

    for _, tc := range tests {
        result := ToUpper(tc.input)
        if result != tc.expected {
            t.Errorf("ToUpper(%q) = %q, want %q",
                tc.input, result, tc.expected)
        }
    }
}
```

### Complex Cases with Multiple Fields
```go
type mathTestCase struct {
    name     string
    a, b     int
    expected int
    hasError bool
}

func TestDivide(t *testing.T) {
    tests := []mathTestCase{
        {"positive numbers", 10, 2, 5, false},
        {"negative numbers", -10, 2, -5, false},
        {"division by zero", 10, 0, 0, true},
        {"zero dividend", 0, 5, 0, false},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result, err := Divide(tt.a, tt.b)

            if tt.hasError {
                if err == nil {
                    t.Errorf("Divide(%d, %d) expected error, got nil",
                        tt.a, tt.b)
                }
            } else {
                if err != nil {
                    t.Errorf("Divide(%d, %d) unexpected error: %v",
                        tt.a, tt.b, err)
                }
                if result != tt.expected {
                    t.Errorf("Divide(%d, %d) = %d, want %d",
                        tt.a, tt.b, result, tt.expected)
                }
            }
        })
    }
}
```

## Anonymous Structs

```go
func TestStringOperations(t *testing.T) {
    tests := []struct {
        name     string
        input    string
        expected string
    }{
        {"empty string", "", ""},
        {"single word", "hello", "hello"},
        {"with spaces", "hello world", "hello world"},
        {"with numbers", "hello123", "hello123"},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := ProcessString(tt.input)
            if result != tt.expected {
                t.Errorf("ProcessString(%q) = %q, want %q",
                    tt.input, result, tt.expected)
            }
        })
    }
}
```

## Testing Different Data Types

### String Tests
```go
func TestStringFunctions(t *testing.T) {
    tests := []struct {
        name     string
        input    string
        expected int
    }{
        {"count vowels", "hello", 2},
        {"count consonants", "world", 4},
        {"empty string", "", 0},
        {"all vowels", "aeiou", 5},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := CountVowels(tt.input)
            if result != tt.expected {
                t.Errorf("CountVowels(%q) = %d, want %d",
                    tt.input, result, tt.expected)
            }
        })
    }
}
```

### Numeric Tests
```go
func TestNumericFunctions(t *testing.T) {
    tests := []struct {
        name     string
        input    int
        expected bool
    }{
        {"even number", 2, true},
        {"odd number", 3, false},
        {"zero", 0, true},
        {"negative even", -2, true},
        {"negative odd", -3, false},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := IsEven(tt.input)
            if result != tt.expected {
                t.Errorf("IsEven(%d) = %t, want %t",
                    tt.input, result, tt.expected)
            }
        })
    }
}
```

## Testing Error Conditions

```go
func TestErrorHandling(t *testing.T) {
    tests := []struct {
        name     string
        input    string
        expected int
        hasError bool
    }{
        {"valid number", "123", 123, false},
        {"zero", "0", 0, false},
        {"negative", "-5", -5, false},
        {"invalid input", "abc", 0, true},
        {"empty string", "", 0, true},
        {"float", "12.34", 0, true},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result, err := ParseInt(tt.input)

            if tt.hasError {
                if err == nil {
                    t.Errorf("ParseInt(%q) expected error, got nil", tt.input)
                }
            } else {
                if err != nil {
                    t.Errorf("ParseInt(%q) unexpected error: %v", tt.input, err)
                }
                if result != tt.expected {
                    t.Errorf("ParseInt(%q) = %d, want %d",
                        tt.input, result, tt.expected)
                }
            }
        })
    }
}
```

## Testing Edge Cases

```go
func TestEdgeCases(t *testing.T) {
    tests := []struct {
        name     string
        input    []int
        expected int
    }{
        {"empty slice", []int{}, 0},
        {"single element", []int{5}, 5},
        {"multiple elements", []int{1, 2, 3, 4, 5}, 15},
        {"negative numbers", []int{-1, -2, -3}, -6},
        {"mixed numbers", []int{-1, 0, 1}, 0},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := Sum(tt.input)
            if result != tt.expected {
                t.Errorf("Sum(%v) = %d, want %d",
                    tt.input, result, tt.expected)
            }
        })
    }
}
```

## Best Practices

1. **Use descriptive names** - Test case names should explain what's being tested
2. **Include edge cases** - Test empty inputs, boundary values, error conditions
3. **Keep test data simple** - Avoid complex setup in test cases
4. **Use consistent structure** - All test cases should have the same fields
5. **Test both success and failure** - Include error cases in your test data
6. **Use subtests** - Wrap each test case in `t.Run()` for better organization

## Example Program

```go
package main

import (
    "testing"
    "strings"
)

func TestStringOperations(t *testing.T) {
    tests := []struct {
        name     string
        input    string
        expected string
    }{
        {"uppercase", "hello", "HELLO"},
        {"empty string", "", ""},
        {"with numbers", "hello123", "HELLO123"},
        {"with spaces", "hello world", "HELLO WORLD"},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := strings.ToUpper(tt.input)
            if result != tt.expected {
                t.Errorf("ToUpper(%q) = %q, want %q",
                    tt.input, result, tt.expected)
            }
        })
    }
}
```

## Next Steps

Now that you understand table-driven tests, try the tasks in `tasks.md`!