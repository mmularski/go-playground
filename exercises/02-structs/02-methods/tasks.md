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

## Example structure for student/main.go:
```go
package main

import "fmt"

// Task 1: Basic structs and methods
type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return 3.14159 * c.Radius * c.Radius
}

func (c Circle) Circumference() float64 {
    return 2 * 3.14159 * c.Radius
}

func (c Circle) IsValid() bool {
    return c.Radius > 0
}

// Task 2: Pointer receivers
type BankAccount struct {
    Balance float64
}

func (b *BankAccount) Deposit(amount float64) {
    b.Balance += amount
}

func (b *BankAccount) Withdraw(amount float64) error {
    if amount > b.Balance {
        return fmt.Errorf("insufficient funds")
    }
    b.Balance -= amount
    return nil
}

func (b BankAccount) GetBalance() float64 {
    return b.Balance
}

// Task 3: Method chaining
type StringBuilder struct {
    content string
}

func (sb *StringBuilder) Append(text string) *StringBuilder {
    sb.content += text
    return sb
}

func (sb *StringBuilder) AppendLine(text string) *StringBuilder {
    sb.content += text + "\n"
    return sb
}

func (sb *StringBuilder) Clear() *StringBuilder {
    sb.content = ""
    return sb
}

func (sb StringBuilder) ToString() string {
    return sb.content
}

// Task 4: Custom types
type Age int

func (a Age) IsAdult() bool {
    return a >= 18
}

func (a Age) String() string {
    return fmt.Sprintf("%d years old", a)
}

type Email string

func (e Email) IsValid() bool {
    // Simple validation - contains @
    return len(e) > 0 && string(e) != ""
}

func (e Email) Domain() string {
    // Extract domain from email
    return ""
}

// Task 5: Interface methods
type Shape interface {
    Area() float64
}

type Rectangle struct {
    Width, Height float64
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func TotalArea(shapes []Shape) float64 {
    total := 0.0
    for _, shape := range shapes {
        total += shape.Area()
    }
    return total
}

func main() {
    // Demonstrate your methods here
}
```

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