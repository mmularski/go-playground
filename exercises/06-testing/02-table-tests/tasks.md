# Table-Driven Tests - Tasks

## Task 1: Basic Table-Driven Test
Create a simple table-driven test.

**Requirements:**
- Create a test function with multiple test cases
- Use a slice of structs to define test data
- Test a simple function with different inputs
- Use descriptive test case names

## Task 2: Testing String Functions
Create table-driven tests for string operations.

**Requirements:**
- Test string manipulation functions
- Include edge cases (empty strings, special characters)
- Use proper string formatting in error messages
- Test both success and edge cases

## Task 3: Testing Numeric Functions
Create table-driven tests for numeric operations.

**Requirements:**
- Test mathematical functions
- Include positive, negative, and zero values
- Test boundary conditions
- Handle different numeric types

## Task 4: Testing Error Conditions
Create table-driven tests that check error handling.

**Requirements:**
- Test functions that can return errors
- Include both valid and invalid inputs
- Check error messages and types
- Test error conditions systematically

## Task 5: Testing Edge Cases
Create table-driven tests for edge cases.

**Requirements:**
- Test boundary values
- Test empty or nil inputs
- Test extreme values
- Ensure comprehensive coverage

## Task 6: Complex Test Cases
Create table-driven tests with complex data structures.

**Requirements:**
- Test functions with multiple parameters
- Use complex input types (slices, maps)
- Test multiple return values
- Demonstrate advanced table-driven patterns

---

## How to complete these tasks:

1. Navigate to the `student/` directory
2. Create your code in `main.go`
3. Create tests in `main_test.go`
4. Run your tests: `go test`

## Example structure for student/main.go:
```go
package main

import (
    "fmt"
    "strconv"
    "strings"
)

// Task 1: Basic string function
func ToUpper(s string) string {
    return strings.ToUpper(s)
}

// Task 2: String processing functions
func CountVowels(s string) int {
    vowels := "aeiouAEIOU"
    count := 0
    for _, char := range s {
        if strings.ContainsRune(vowels, char) {
            count++
        }
    }
    return count
}

func ReverseString(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

// Task 3: Numeric functions
func IsEven(n int) bool {
    return n%2 == 0
}

func Abs(n int) int {
    if n < 0 {
        return -n
    }
    return n
}

func Sum(numbers []int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}

// Task 4: Functions with error handling
func ParseInt(s string) (int, error) {
    return strconv.Atoi(s)
}

func Divide(a, b int) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("division by zero")
    }
    return float64(a) / float64(b), nil
}

// Task 5: Edge case functions
func FindMax(numbers []int) (int, error) {
    if len(numbers) == 0 {
        return 0, fmt.Errorf("empty slice")
    }
    max := numbers[0]
    for _, num := range numbers[1:] {
        if num > max {
            max = num
        }
    }
    return max, nil
}

func IsPalindrome(s string) bool {
    if len(s) <= 1 {
        return true
    }
    return s[0] == s[len(s)-1] && IsPalindrome(s[1:len(s)-1])
}

// Task 6: Complex functions
type Person struct {
    Name string
    Age  int
}

func FilterAdults(people []Person) []Person {
    var adults []Person
    for _, person := range people {
        if person.Age >= 18 {
            adults = append(adults, person)
        }
    }
    return adults
}

func GroupByAge(people []Person) map[int][]Person {
    groups := make(map[int][]Person)
    for _, person := range people {
        groups[person.Age] = append(groups[person.Age], person)
    }
    return groups
}
```

## Example structure for student/main_test.go:
```go
package main

import (
    "testing"
)

// Task 1: Basic table-driven test
func TestToUpper(t *testing.T) {
    tests := []struct {
        name     string
        input    string
        expected string
    }{
        {"empty string", "", ""},
        {"single word", "hello", "HELLO"},
        {"multiple words", "hello world", "HELLO WORLD"},
        {"with numbers", "hello123", "HELLO123"},
        {"already uppercase", "HELLO", "HELLO"},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := ToUpper(tt.input)
            if result != tt.expected {
                t.Errorf("ToUpper(%q) = %q, want %q",
                    tt.input, result, tt.expected)
            }
        })
    }
}

// Task 2: String function tests
func TestCountVowels(t *testing.T) {
    tests := []struct {
        name     string
        input    string
        expected int
    }{
        {"no vowels", "xyz", 0},
        {"all vowels", "aeiou", 5},
        {"mixed case", "Hello World", 3},
        {"empty string", "", 0},
        {"numbers only", "12345", 0},
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

func TestReverseString(t *testing.T) {
    tests := []struct {
        name     string
        input    string
        expected string
    }{
        {"empty string", "", ""},
        {"single character", "a", "a"},
        {"palindrome", "racecar", "racecar"},
        {"normal word", "hello", "olleh"},
        {"with spaces", "hello world", "dlrow olleh"},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := ReverseString(tt.input)
            if result != tt.expected {
                t.Errorf("ReverseString(%q) = %q, want %q",
                    tt.input, result, tt.expected)
            }
        })
    }
}

// Task 3: Numeric function tests
func TestIsEven(t *testing.T) {
    tests := []struct {
        name     string
        input    int
        expected bool
    }{
        {"positive even", 2, true},
        {"positive odd", 3, false},
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

func TestAbs(t *testing.T) {
    tests := []struct {
        name     string
        input    int
        expected int
    }{
        {"positive number", 5, 5},
        {"negative number", -5, 5},
        {"zero", 0, 0},
        {"large positive", 1000, 1000},
        {"large negative", -1000, 1000},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := Abs(tt.input)
            if result != tt.expected {
                t.Errorf("Abs(%d) = %d, want %d",
                    tt.input, result, tt.expected)
            }
        })
    }
}

// Task 4: Error handling tests
func TestParseInt(t *testing.T) {
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

// Task 5: Edge case tests
func TestFindMax(t *testing.T) {
    tests := []struct {
        name     string
        input    []int
        expected int
        hasError bool
    }{
        {"single element", []int{5}, 5, false},
        {"multiple elements", []int{1, 5, 3, 9, 2}, 9, false},
        {"negative numbers", []int{-1, -5, -3}, -1, false},
        {"empty slice", []int{}, 0, true},
        {"all same", []int{5, 5, 5}, 5, false},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result, err := FindMax(tt.input)

            if tt.hasError {
                if err == nil {
                    t.Errorf("FindMax(%v) expected error, got nil", tt.input)
                }
            } else {
                if err != nil {
                    t.Errorf("FindMax(%v) unexpected error: %v", tt.input, err)
                }
                if result != tt.expected {
                    t.Errorf("FindMax(%v) = %d, want %d",
                        tt.input, result, tt.expected)
                }
            }
        })
    }
}

// Task 6: Complex test cases
func TestFilterAdults(t *testing.T) {
    tests := []struct {
        name     string
        input    []Person
        expected []Person
    }{
        {"all adults", []Person{{"Alice", 25}, {"Bob", 30}},
         []Person{{"Alice", 25}, {"Bob", 30}}},
        {"mixed ages", []Person{{"Alice", 25}, {"Child", 10}, {"Bob", 30}},
         []Person{{"Alice", 25}, {"Bob", 30}}},
        {"no adults", []Person{{"Child1", 10}, {"Child2", 15}},
         []Person{}},
        {"empty list", []Person{}, []Person{}},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := FilterAdults(tt.input)
            if len(result) != len(tt.expected) {
                t.Errorf("FilterAdults(%v) returned %d people, want %d",
                    tt.input, len(result), len(tt.expected))
            }
        })
    }
}
```

## Testing Your Code

Run your tests:
```bash
cd student/
go test
```

Run with verbose output:
```bash
go test -v
```

You should see output demonstrating all table-driven test concepts.