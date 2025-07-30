# Basic Structs - Tasks

## Task 1: Create Simple Structs
Create basic structs with different field types.

**Requirements:**
- Create a struct for a Person with name, age, and email
- Create a struct for a Rectangle with width and height
- Create a struct for a Book with title, author, and year
- Initialize structs using different methods
- Print struct values

## Task 2: Struct Fields and Tags
Work with struct fields and tags.

**Requirements:**
- Create structs with different field types (int, string, bool, float64)
- Use field tags for JSON serialization
- Access struct fields using dot notation
- Modify struct fields
- Demonstrate field visibility (exported vs unexported)

## Task 3: Struct Initialization
Practice different ways to initialize structs.

**Requirements:**
- Initialize structs using struct literals
- Initialize structs using field names
- Initialize structs using positional values
- Create structs with default values
- Use the new() function to create structs

## Task 4: Nested Structs
Work with structs containing other structs.

**Requirements:**
- Create nested structs (structs within structs)
- Access nested struct fields
- Initialize nested structs
- Create structs with slice and map fields
- Demonstrate struct composition

## Task 5: Anonymous Structs
Create and use anonymous structs.

**Requirements:**
- Create anonymous structs for one-time use
- Use anonymous structs in function parameters
- Create anonymous structs with embedded structs
- Demonstrate when to use anonymous vs named structs

## Task 6: Struct Methods
Create methods for your structs.

**Requirements:**
- Create methods with value receivers
- Create methods with pointer receivers
- Demonstrate the difference between value and pointer receivers
- Create getter and setter methods
- Create methods that modify struct state

---

## How to complete these tasks:

1. Navigate to the `student/` directory
2. Create your code in `main.go`
3. Run your code: `go run main.go`

## Example structure for student/main.go:
```go
package main

import "fmt"

// Task 1: Basic structs
type Person struct {
    Name  string
    Age   int
    Email string
}

type Rectangle struct {
    Width  float64
    Height float64
}

type Book struct {
    Title  string
    Author string
    Year   int
}

func main() {
    // Task 1: Create and use structs
    person := Person{Name: "Alice", Age: 25, Email: "alice@example.com"}
    fmt.Printf("Person: %+v\n", person)

    // Task 2: Struct fields and tags
    rect := Rectangle{Width: 10.5, Height: 5.2}
    fmt.Printf("Rectangle: %+v\n", rect)

    // Task 3: Different initialization methods
    book := Book{"Go Programming", "John Doe", 2023}
    fmt.Printf("Book: %+v\n", book)

    // Task 4: Nested structs
    address := Address{City: "New York", Country: "USA"}
    personWithAddress := PersonWithAddress{
        Person:  person,
        Address: address,
    }
    fmt.Printf("Person with address: %+v\n", personWithAddress)

    // Task 5: Anonymous structs
    config := struct {
        Port     int
        Hostname string
        Debug    bool
    }{
        Port:     8080,
        Hostname: "localhost",
        Debug:    true,
    }
    fmt.Printf("Config: %+v\n", config)

    // Task 6: Struct methods
    area := rect.Area()
    fmt.Printf("Rectangle area: %.2f\n", area)

    isAdult := person.IsAdult()
    fmt.Printf("Is adult: %t\n", isAdult)
}

// Task 4: Nested structs
type Address struct {
    City    string
    Country string
}

type PersonWithAddress struct {
    Person  Person
    Address Address
}

// Task 6: Methods
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func (p Person) IsAdult() bool {
    return p.Age >= 18
}
```

## Testing Your Code

Run your program:
```bash
cd student/
go run main.go
```

You should see output demonstrating all struct concepts.