# Interfaces - Overview

## What are Interfaces?

Interfaces define a set of method signatures that types must implement. They provide a way to achieve polymorphism and abstraction in Go.

## Interface Declaration

### Basic Interface
```go
type Shape interface {
    Area() float64
    Perimeter() float64
    String() string
}
```

### Interface with Multiple Methods
```go
type Animal interface {
    Speak() string
    Move() string
}
```

## Implementing Interfaces

A type implements an interface by implementing all its methods:

```go
type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
    return 2 * math.Pi * c.Radius
}

func (c Circle) String() string {
    return fmt.Sprintf("Circle{Radius: %.2f}", c.Radius)
}

// Circle now implements Shape interface
```

## Interface Composition

Interfaces can embed other interfaces:

```go
type ColoredShape interface {
    Shape
    Color() string
}

type ColoredCircle struct {
    Circle
    color string
}

func (cc ColoredCircle) Color() string {
    return cc.color
}

// ColoredCircle implements ColoredShape interface
```

## Empty Interface

The empty interface `interface{}` accepts any type:

```go
func printAnything(v interface{}) {
    fmt.Printf("Value: %v, Type: %T\n", v, v)
}

// Usage:
printAnything(42)
printAnything("Hello")
printAnything(Circle{Radius: 1})
```

## Type Assertions

Check if an interface value holds a specific type:

```go
func processShape(s Shape) {
    // Type assertion with ok check
    if circle, ok := s.(Circle); ok {
        fmt.Printf("This is a circle with radius %.2f\n", circle.Radius)
    } else if rect, ok := s.(Rectangle); ok {
        fmt.Printf("This is a rectangle %.2f x %.2f\n", rect.Width, rect.Height)
    }
}
```

## Type Switches

Use switch statements with type assertions:

```go
func describeShape(s Shape) {
    switch shape := s.(type) {
    case Circle:
        fmt.Printf("Circle with radius %.2f\n", shape.Radius)
    case Rectangle:
        fmt.Printf("Rectangle %.2f x %.2f\n", shape.Width, shape.Height)
    default:
        fmt.Printf("Unknown shape: %T\n", shape)
    }
}
```

## Interface Best Practices

### 1. Keep Interfaces Small
```go
// Good - small interface
type Reader interface {
    Read(p []byte) (n int, err error)
}

// Good - small interface
type Writer interface {
    Write(p []byte) (n int, err error)
}

// Good - composition
type ReadWriter interface {
    Reader
    Writer
}
```

### 2. Define Interfaces Where They Are Used
```go
// Define interface in the package that uses it
func calculateTotalArea(shapes []Shape) float64 {
    total := 0.0
    for _, shape := range shapes {
        total += shape.Area()
    }
    return total
}
```

### 3. Use Interface{} Sparingly
```go
// Prefer specific interfaces
func processData(data []byte) error {
    // Better than interface{}
}

// Only use interface{} when you truly need to accept any type
func logValue(v interface{}) {
    fmt.Printf("Value: %v\n", v)
}
```

## Common Interface Patterns

### 1. Interface as Function Parameter
```go
func makeAnimalSpeak(a Animal) {
    fmt.Printf("Animal says: %s\n", a.Speak())
}

// Usage:
dog := Dog{Name: "Buddy"}
makeAnimalSpeak(dog)
```

### 2. Interface Slices
```go
shapes := []Shape{
    Circle{Radius: 3},
    Rectangle{Width: 4, Height: 5},
}

for _, shape := range shapes {
    fmt.Printf("Area: %.2f\n", shape.Area())
}
```

### 3. Interface with State
```go
type Counter interface {
    Increment()
    GetCount() int
}

type SimpleCounter struct {
    count int
}

func (sc *SimpleCounter) Increment() {
    sc.count++
}

func (sc SimpleCounter) GetCount() int {
    return sc.count
}
```

## Interface Nil Check

```go
var shape Shape
if shape == nil {
    fmt.Println("Shape is nil")
}
```

## Example Program

```go
package main

import "fmt"

type Shape interface {
    Area() float64
    String() string
}

type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return 3.14159 * c.Radius * c.Radius
}

func (c Circle) String() string {
    return fmt.Sprintf("Circle{Radius: %.2f}", c.Radius)
}

type Rectangle struct {
    Width  float64
    Height float64
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func (r Rectangle) String() string {
    return fmt.Sprintf("Rectangle{Width: %.2f, Height: %.2f}", r.Width, r.Height)
}

func calculateTotalArea(shapes []Shape) float64 {
    total := 0.0
    for _, shape := range shapes {
        total += shape.Area()
    }
    return total
}

func main() {
    shapes := []Shape{
        Circle{Radius: 3},
        Rectangle{Width: 4, Height: 5},
    }

    totalArea := calculateTotalArea(shapes)
    fmt.Printf("Total area: %.2f\n", totalArea)
}
```

## Next Steps

Now that you understand interfaces, try the tasks in `tasks.md`!
