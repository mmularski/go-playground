# Basic Structs - Overview

## What are Structs?

Structs are composite data types that group together variables under a single name. They allow you to create custom data types that represent real-world objects.

## Struct Declaration

### Basic Struct
```go
type Person struct {
    Name string
    Age  int
    City string
}
```

### Struct with Different Field Types
```go
type Student struct {
    Person     // embedded struct
    StudentID  string
    GPA        float64
    IsEnrolled bool
    Courses    []string
    Grades     map[string]int
}
```

## Creating Struct Instances

### Method 1: Field Names
```go
person := Person{
    Name: "John",
    Age:  30,
    City: "New York",
}
```

### Method 2: Positional (Order Matters)
```go
person := Person{"Jane", 25, "Los Angeles"}
```

### Method 3: Zero Value
```go
var person Person
person.Name = "Bob"
person.Age = 35
person.City = "Chicago"
```

## Struct Tags

Tags provide metadata for fields, commonly used for JSON serialization:

```go
type Product struct {
    ID          int     `json:"id" db:"product_id"`
    Name        string  `json:"name" db:"product_name"`
    Price       float64 `json:"price" db:"price"`
    Description string  `json:"description,omitempty"`
    InStock     bool    `json:"in_stock"`
}
```

## Anonymous Structs

Structs without a name, created inline:

```go
car := struct {
    Brand string
    Model string
    Year  int
}{
    Brand: "Toyota",
    Model: "Camry",
    Year:  2020,
}
```

## Accessing Struct Fields

```go
person := Person{"John", 30, "NYC"}

// Access fields
name := person.Name
age := person.Age

// Modify fields
person.Age = 31
person.City = "Boston"
```

## Embedded Structs

Structs can embed other structs:

```go
type Person struct {
    Name string
    Age  int
}

type Student struct {
    Person     // embedded
    StudentID  string
    GPA        float64
}

// Access embedded fields
student := Student{
    Person: Person{"Alice", 20},
    StudentID: "STU001",
    GPA: 3.8,
}

fmt.Println(student.Name) // Access embedded field
```

## Struct Methods

Structs can have methods with value or pointer receivers:

```go
type Counter struct {
    count int
}

// Pointer receiver (can modify struct)
func (c *Counter) Increment() {
    c.count++
}

// Value receiver (cannot modify struct)
func (c Counter) GetCount() int {
    return c.count
}
```

## Struct Comparison

Structs can be compared if all fields are comparable:

```go
p1 := Person{"John", 30, "NYC"}
p2 := Person{"John", 30, "NYC"}
p3 := Person{"Jane", 30, "NYC"}

fmt.Println(p1 == p2) // true
fmt.Println(p1 == p3) // false
```

## Struct Pointers

```go
person := &Person{"Mike", 28, "Seattle"}
fmt.Printf("Person: %+v\n", *person)

// Modify via pointer
person.Age = 29
```

## Best Practices

1. **Use descriptive field names**: `FirstName` instead of `fn`
2. **Group related fields** together
3. **Use tags for serialization** when needed
4. **Choose appropriate receivers** for methods
5. **Use embedded structs** for composition

## Example Program

```go
package main

import "fmt"

type Person struct {
    Name string
    Age  int
    City string
}

type Student struct {
    Person
    StudentID string
    GPA       float64
}

func main() {
    // Create a person
    person := Person{
        Name: "John",
        Age:  30,
        City: "New York",
    }

    // Create a student
    student := Student{
        Person:    person,
        StudentID: "STU001",
        GPA:       3.8,
    }

    fmt.Printf("Student: %+v\n", student)
    fmt.Printf("Name: %s, GPA: %.1f\n", student.Name, student.GPA)
}
```

## Next Steps

Now that you understand basic structs, try the tasks in `tasks.md`!
