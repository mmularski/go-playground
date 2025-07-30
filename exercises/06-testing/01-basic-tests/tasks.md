# Basic Testing - Tasks

## Task 1: Create Your First Test
Create a simple function and write a test for it.

**Requirements:**
- Create a function `Add(a, b int) int` that adds two integers
- Write a test function `TestAdd` that verifies the function works correctly
- Test with at least 3 different input combinations
- Run the test and verify it passes

## Task 2: Test Error Handling
Create a function that can return an error and test both success and error cases.

**Requirements:**
- Create a function `Divide(a, b int) (float64, error)` that divides two numbers
- Handle division by zero error
- Write tests for successful division
- Write tests for division by zero error
- Verify error message is correct

## Task 3: Test String Functions
Create and test functions that work with strings.

**Requirements:**
- Create a function `Reverse(s string) string` that reverses a string
- Create a function `CountVowels(s string) int` that counts vowels in a string
- Write comprehensive tests for both functions
- Test edge cases (empty string, single character, etc.)

## Task 4: Test Boolean Functions
Create and test functions that return boolean values.

**Requirements:**
- Create a function `IsPalindrome(s string) bool` that checks if a string is a palindrome
- Create a function `IsPrime(n int) bool` that checks if a number is prime
- Write tests for both functions
- Test with various inputs including edge cases

## Task 5: Test Slice Functions
Create and test functions that work with slices.

**Requirements:**
- Create a function `SumSlice(numbers []int) int` that sums all numbers in a slice
- Create a function `FindMax(numbers []int) (int, error)` that finds the maximum value
- Handle empty slice case in FindMax
- Write comprehensive tests for both functions

## Task 6: Helper Functions
Create helper functions to make your tests more readable and maintainable.

**Requirements:**
- Create a helper function `assertEqual(t *testing.T, got, want int)` for comparing integers
- Create a helper function `assertStringEqual(t *testing.T, got, want string)` for comparing strings
- Create a helper function `assertError(t *testing.T, err error, expectedMsg string)` for checking errors
- Refactor your existing tests to use these helpers

## Task 7: Test Organization
Organize your tests using table-driven tests and subtests.

**Requirements:**
- Refactor your tests to use table-driven tests where appropriate
- Use `t.Run()` to create subtests for better organization
- Group related test cases together
- Make your tests more readable and maintainable

## Task 8: Test Coverage
Ensure your tests have good coverage.

**Requirements:**
- Run `go test -cover` to see your test coverage
- Aim for at least 80% coverage
- Add tests for any uncovered code paths
- Run `go test -coverprofile=coverage.out` and generate HTML report

---

## How to complete these tasks:

1. Navigate to the `student/` directory
2. Create your functions in `main.go`
3. Create your tests in `main_test.go`
4. Run tests: `go test`
5. Run tests with coverage: `go test -cover`


## Example structure for student/main_test.go:
```go
package main

import "testing"

// Task 1: Test Add function
func TestAdd(t *testing.T) {
    // TODO: Write tests for Add function
}

// Task 2: Test Divide function
func TestDivide(t *testing.T) {
    // TODO: Write tests for Divide function
}

// Task 3: Test string functions
func TestReverse(t *testing.T) {
    // TODO: Write tests for Reverse function
}

func TestCountVowels(t *testing.T) {
    // TODO: Write tests for CountVowels function
}

// Task 4: Test boolean functions
func TestIsPalindrome(t *testing.T) {
    // TODO: Write tests for IsPalindrome function
}

func TestIsPrime(t *testing.T) {
    // TODO: Write tests for IsPrime function
}

// Task 5: Test slice functions
func TestSumSlice(t *testing.T) {
    // TODO: Write tests for SumSlice function
}

func TestFindMax(t *testing.T) {
    // TODO: Write tests for FindMax function
}

// Task 6: Helper functions
func assertEqual(t *testing.T, got, want int) {
    // TODO: Implement helper function
}

func assertStringEqual(t *testing.T, got, want string) {
    // TODO: Implement helper function
}

func assertError(t *testing.T, err error, expectedMsg string) {
    // TODO: Implement helper function
}
```

## Testing Your Code

Run your tests:
```bash
cd student/
go test
```

Run tests with verbose output:
```bash
go test -v
```

Run tests with coverage:
```bash
go test -cover
```

Generate coverage report:
```bash
go test -coverprofile=coverage.out
go tool cover -html=coverage.out
```

You should see output like:
```
=== RUN   TestAdd
--- PASS: TestAdd (0.00s)
=== RUN   TestDivide
--- PASS: TestDivide (0.00s)
PASS
coverage: 85.7% of statements
ok      ./student   0.001s
```