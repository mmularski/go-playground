# Methods - Tasks

## Task 1: Basic Methods
Create methods with value receivers.

**Requirements:**
- Create a struct for a Circle with radius
- Create a method `Area()` that calculates circle area
- Create a method `Circumference()` that calculates circumference
- Create a method `IsValid()` that checks if radius is positive
- Test your methods

## Task 2: Pointer Receivers
Create methods that modify struct state.

**Requirements:**
- Create a struct for a BankAccount with balance
- Create a method `Deposit(amount float64)` that adds to balance
- Create a method `Withdraw(amount float64) error` that subtracts from balance
- Create a method `GetBalance() float64` that returns current balance
- Handle insufficient funds in Withdraw

## Task 3: Method Chaining
Create methods that return the receiver for chaining.

**Requirements:**
- Create a struct for a StringBuilder with content
- Create a method `Append(text string) *StringBuilder`
- Create a method `AppendLine(text string) *StringBuilder`
- Create a method `Clear() *StringBuilder`
- Create a method `ToString() string`
- Demonstrate method chaining

## Task 4: Methods on Custom Types
Create methods for custom type definitions.

**Requirements:**
- Create a custom type `type Age int`
- Create a method `IsAdult() bool` for Age
- Create a method `String() string` for Age
- Create a custom type `type Email string`
- Create a method `IsValid() bool` for Email
- Create a method `Domain() string` for Email

## Task 5: Methods with Interfaces
Create methods that work with interfaces.

**Requirements:**
- Create an interface `Shape` with `Area() float64` method
- Create structs `Circle` and `Rectangle` that implement Shape
- Create a method `TotalArea(shapes []Shape) float64`
- Create a method `LargestShape(shapes []Shape) Shape`
- Demonstrate polymorphism

## Task 6: Testing Your Methods
Create functions and tests for your methods.

**Requirements:**
- Create a function `CreateCircle(radius float64) Circle`
- Create a function `CreateBankAccount(initialBalance float64) BankAccount`
- Create a function `ProcessTransaction(account *BankAccount, amount float64) error`
- Create a function `ValidateEmail(email Email) bool`
- Write comprehensive tests for all functions

---

## How to complete these tasks:

1. Navigate to the `student/` directory
2. Create your code in `main.go`
3. Create your tests in `main_test.go`
4. Run your code: `go run main.go`
5. Run your tests: `go test`


## Example structure for student/main_test.go:
```go
package main

import "testing"

func TestCreateCircle(t *testing.T) {
    // TODO: Write tests for CreateCircle function
}

func TestCreateBankAccount(t *testing.T) {
    // TODO: Write tests for CreateBankAccount function
}

func TestProcessTransaction(t *testing.T) {
    // TODO: Write tests for ProcessTransaction function
}

func TestValidateEmail(t *testing.T) {
    // TODO: Write tests for ValidateEmail function
}
```

## Testing Your Code

Run your program:
```bash
cd student/
go run main.go
```

Run your tests:
```bash
go test
```

Run tests with verbose output:
```bash
go test -v
```

You should see output demonstrating all method concepts and comprehensive test coverage.