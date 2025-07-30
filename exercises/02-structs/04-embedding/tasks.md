# Struct Embedding - Tasks

## Task 1: Basic Embedding
Create embedded structs and access their fields.

**Requirements:**
- Create a base struct `Person` with name and age
- Create a struct `Employee` that embeds `Person`
- Add employee-specific fields like salary and department
- Access both embedded and new fields

## Task 2: Method Promotion
Create methods on embedded structs and see them promoted.

**Requirements:**
- Create a method on the base struct
- Create a struct that embeds the base struct
- Demonstrate that the embedded struct gets the method
- Show method promotion in action

## Task 3: Method Overriding
Override methods from embedded structs.

**Requirements:**
- Create a method on the embedded struct
- Create the same method on the embedding struct
- Show how the embedding struct's method takes precedence
- Demonstrate method overriding

## Task 4: Multiple Embedding
Embed multiple structs in one struct.

**Requirements:**
- Create multiple base structs
- Create a struct that embeds multiple other structs
- Access fields from all embedded structs
- Handle potential field name conflicts

## Task 5: Interface Embedding
Embed interfaces to create new interfaces.

**Requirements:**
- Create basic interfaces
- Create a new interface that embeds multiple interfaces
- Implement the composed interface
- Demonstrate interface composition

## Task 6: Practical Application
Create a simple vehicle hierarchy using embedding.

**Requirements:**
- Create a base `Vehicle` struct with common fields
- Create `Car` and `Motorcycle` that embed `Vehicle`
- Add vehicle-specific methods
- Demonstrate polymorphism through embedding

---

## How to complete these tasks:

1. Navigate to the `student/` directory
2. Create your code in `main.go`
3. Run your code: `go run main.go`

## Example structure for student/main.go:
```go
package main

import "fmt"

// Task 1: Basic embedding
type Person struct {
    Name string
    Age  int
}

type Employee struct {
    Person
    Salary     float64
    Department string
}

// Task 2: Method promotion
func (p Person) Introduce() string {
    return fmt.Sprintf("Hi, I'm %s and I'm %d years old", p.Name, p.Age)
}

// Task 3: Method overriding
func (e Employee) Introduce() string {
    return fmt.Sprintf("Hi, I'm %s, I work in %s department", e.Name, e.Department)
}

// Task 4: Multiple embedding
type Address struct {
    Street  string
    City    string
    Country string
}

type Contact struct {
    Email string
    Phone string
}

type Customer struct {
    Person
    Address
    Contact
}

// Task 5: Interface embedding
type Reader interface {
    Read() string
}

type Writer interface {
    Write(string)
}

type ReadWriter interface {
    Reader
    Writer
}

type File struct {
    content string
}

func (f *File) Read() string {
    return f.content
}

func (f *File) Write(data string) {
    f.content = data
}

// Task 6: Practical application
type Vehicle struct {
    Brand string
    Model string
    Year  int
}

func (v Vehicle) GetInfo() string {
    return fmt.Sprintf("%s %s (%d)", v.Brand, v.Model, v.Year)
}

type Car struct {
    Vehicle
    Doors int
}

func (c Car) GetInfo() string {
    return fmt.Sprintf("%s with %d doors", c.Vehicle.GetInfo(), c.Doors)
}

type Motorcycle struct {
    Vehicle
    EngineSize int
}

func (m Motorcycle) GetInfo() string {
    return fmt.Sprintf("%s with %dcc engine", m.Vehicle.GetInfo(), m.EngineSize)
}

func main() {
    // Demonstrate your embedding here
}
```

## Testing Your Code

Run your program:
```bash
cd student/
go run main.go
```

You should see output demonstrating all embedding concepts.