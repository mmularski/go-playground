# Methods - Overview

## What are Methods?

Methods are functions that are associated with a specific type. They allow you to define behavior for your custom types, making your code more organized and object-oriented.

## Method Declaration

### Basic Method Syntax
```go
func (receiver Type) MethodName(parameters) returnType {
    // method body
}
```

## Value Receivers vs Pointer Receivers

### Value Receiver (Creates a Copy)
```go
type Rectangle struct {
    Width  float64
    Height float64
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
    return 2 * (r.Width + r.Height)
}
```

### Pointer Receiver (Modifies Original)
```go
type BankAccount struct {
    Balance float64
    Owner   string
}

func (b *BankAccount) Deposit(amount float64) {
    if amount > 0 {
        b.Balance += amount
    }
}

func (b *BankAccount) Withdraw(amount float64) bool {
    if amount > 0 && b.Balance >= amount {
        b.Balance -= amount
        return true
    }
    return false
}
```

## When to Use Each Type

### Use Value Receivers When:
- The method doesn't need to modify the struct
- The struct is small and copying is efficient
- You want to ensure the original data isn't changed

### Use Pointer Receivers When:
- The method needs to modify the struct
- The struct is large and copying would be expensive
- You want to share the same data across method calls

## Methods on Custom Types

You can define methods on any custom type:

```go
type Temperature float64

func (t Temperature) Celsius() float64 {
    return float64(t)
}

func (t Temperature) Fahrenheit() float64 {
    return float64(t)*9/5 + 32
}

func (t Temperature) String() string {
    return fmt.Sprintf("%.1f°C", t.Celsius())
}
```

## Method Chaining

Methods can return the receiver to enable chaining:

```go
type Calculator struct {
    result float64
}

func (c *Calculator) Add(x float64) *Calculator {
    c.result += x
    return c
}

func (c *Calculator) Multiply(x float64) *Calculator {
    c.result *= x
    return c
}

func (c Calculator) GetResult() float64 {
    return c.result
}

// Usage:
calc := &Calculator{}
result := calc.Add(5).Multiply(2).GetResult() // 10
```

## String Method

The `String()` method is special - it's used by `fmt` functions:

```go
func (r Rectangle) String() string {
    return fmt.Sprintf("Rectangle{Width: %.2f, Height: %.2f}", r.Width, r.Height)
}

// Now you can use:
rect := Rectangle{5, 3}
fmt.Println(rect) // Prints: Rectangle{Width: 5.00, Height: 3.00}
```

## Interface Implementation

Methods allow types to implement interfaces:

```go
type Shape interface {
    Area() float64
    String() string
}

type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return math.Pi * c.Radius * c.Radius
}

func (c Circle) String() string {
    return fmt.Sprintf("Circle{Radius: %.2f}", c.Radius)
}

// Circle now implements Shape interface
```

## Best Practices

1. **Use value receivers** for small structs that don't need modification
2. **Use pointer receivers** for methods that modify the struct
3. **Be consistent** - use the same receiver type for all methods of a struct
4. **Implement String()** for custom types to improve debugging
5. **Use method chaining** when it makes sense for your API

## Example Program

```go
package main

import "fmt"

type Rectangle struct {
    Width  float64
    Height float64
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
    return 2 * (r.Width + r.Height)
}

func (r Rectangle) String() string {
    return fmt.Sprintf("Rectangle{Width: %.2f, Height: %.2f}", r.Width, r.Height)
}

type BankAccount struct {
    Balance float64
    Owner   string
}

func (b *BankAccount) Deposit(amount float64) {
    if amount > 0 {
        b.Balance += amount
    }
}

func (b BankAccount) GetBalance() float64 {
    return b.Balance
}

func main() {
    // Value receiver example
    rect := Rectangle{5, 3}
    fmt.Printf("Area: %.2f\n", rect.Area())

    // Pointer receiver example
    account := &BankAccount{Balance: 1000, Owner: "John"}
    account.Deposit(500)
    fmt.Printf("Balance: %.2f\n", account.GetBalance())
}
```

## Next Steps

Now that you understand methods, try the tasks in `tasks.md`!
