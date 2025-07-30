# Interfaces - Tasks

## Task 1: Basic Interfaces
Create and implement basic interfaces.

**Requirements:**
- Create an interface `Shape` with `Area() float64` method
- Create structs `Circle` and `Rectangle` that implement Shape
- Create a function that works with any Shape
- Demonstrate polymorphism

## Task 2: Interface Composition
Create interfaces that compose other interfaces.

**Requirements:**
- Create interface `Readable` with `Read() string` method
- Create interface `Writable` with `Write(string)` method
- Create interface `ReadWritable` that embeds both
- Implement these interfaces with concrete types

## Task 3: Empty Interface
Work with the empty interface `interface{}`.

**Requirements:**
- Create a function that accepts `interface{}`
- Use type assertions to work with different types
- Use type switches to handle different types
- Demonstrate when to use empty interface

## Task 4: Interface Implementation
Create types that implement multiple interfaces.

**Requirements:**
- Create interface `Stringer` with `String() string` method
- Create interface `Comparable` with `Compare(other T) int` method
- Create a type that implements both interfaces
- Use the type in different contexts

## Task 5: Interface Best Practices
Follow Go interface design principles.

**Requirements:**
- Keep interfaces small (1-3 methods)
- Define interfaces where they are used
- Use interfaces for behavior, not data
- Create interfaces that are easy to implement

## Task 6: Practical Application
Create a simple drawing system using interfaces.

**Requirements:**
- Create interface `Drawable` with `Draw()` method
- Create different shapes that implement Drawable
- Create a canvas that can draw any Drawable
- Demonstrate the power of interfaces

---

## How to complete these tasks:

1. Navigate to the `student/` directory
2. Create your code in `main.go`
3. Run your code: `go run main.go`


## Testing Your Code

Run your program:
```bash
cd student/
go run main.go
```

You should see output demonstrating all interface concepts.